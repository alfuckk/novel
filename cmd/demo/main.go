package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {

	// Collector for index page
	// cIndex := colly.NewCollector()
	cIndex := colly.NewCollector(
		// colly.MaxDepth(5), // URL的最大深度
		colly.Async(true), // 启用异步模式
	)
	// Visit index page
	cIndex.OnHTML(".nav a", func(e *colly.HTMLElement) {
		category := e.Text
		link := e.Attr("href")
		fmt.Printf("Found category: %s at %s\n", category, link)

		// Collector for list page
		cList := cIndex.Clone()
		SearchList(cList)

		// Visit list page

		fmt.Println("列表页：", "https://quanben.io"+link)
		cList.Visit("https://quanben.io" + link)

	})

	cIndex.Visit("https://quanben.io")

	fmt.Println("Done!")
}

func SearchList(c *colly.Collector) {
	c.OnHTML(".box .row a", func(e *colly.HTMLElement) {
		name := e.Text
		detailURL := e.Attr("href")
		fmt.Printf("- Found novel: %s at %s\n", name, "https://quanben.io"+detailURL)
		// Collector for detail page
		cDetail := c.Clone()
		// Visit detail page
		cDetail.OnHTML(".box", func(e *colly.HTMLElement) {
			chapters := e.ChildAttr(".tool_button a", "href")
			fmt.Printf("-- Got chapters: %s\n", "https://quanben.io"+chapters)
			chapterList := c.Clone()
			chapterList.OnHTML(".box", func(e *colly.HTMLElement) {
				category := e.ChildText(".list2 p span")
				fmt.Printf("-- Got category: %s\n", category)
				title := e.ChildText(".list2 h3 span")
				// fmt.Println(title)
				fmt.Printf("-- Got chapter list: %s\n", title)
				contentURLs := e.ChildAttrs(".list3 a", "href")
				if len(contentURLs) > 0 {
					re := regexp.MustCompile("[0-9]+")
					numbers := re.FindAllString(contentURLs[len(contentURLs)-1], -1)
					total, _ := strconv.Atoi(numbers[0])
					ch := make(chan string, total)
					for i := 1; i < total; i++ {
						contentURL := fmt.Sprintf("https://quanben.io%s%d.html", detailURL, i)
						fmt.Println("URL:", contentURL)
						ch <- contentURL
						// GetContent(contentURL, c)
						// return
					}
					go func() {
						for {
							select {
							case contentURL := <-ch:
								GetContent(contentURL)
							case <-time.After(5 * time.Second):
								fmt.Println("timeout")
								return
							}
						}
					}()
					// fmt.Println("最后一章：", numbers[0])
					// fmt.Printf("- novel: %s at %s\n", name, "https://quanben.io"+detailURL)

					// GetContent(contentURLs[len(contentURLs)-1], c)
				}
			})
			chapterList.Visit("https://quanben.io" + chapters)
		})

		cDetail.Visit("https://quanben.io" + detailURL)
	})
	c.OnHTML(".page_next", func(e *colly.HTMLElement) {
		exists := e.DOM.Find("a").Length() > 0
		if exists {
			cNext := c.Clone()
			SearchList(cNext)
			nextURL := e.ChildAttr("a", "href")
			fmt.Printf("-- Visiting next page %s\n", "https://quanben.io"+nextURL)
			cNext.Visit("https://quanben.io" + nextURL)
		} else {
			fmt.Println("Done!")
		}
	})
}

func GetContent(nextURL string) {
	c := colly.NewCollector(colly.MaxDepth(4))
	time.Sleep(1 * time.Second)
	// content := c.Clone()
	// fmt.Println(content)
	c.OnHTML("#content", func(e *colly.HTMLElement) {
		fmt.Println("当前URL：", nextURL)
		// mulu := e.ChildAttrs(".list_page a", "href")
		title := e.ChildText(".headline")
		fmt.Println("章节名", title)
		// newURL := "https://quanben.io" + mulu[len(mulu)-1]
		// if strings.Contains(newURL, "list") {
		// 	fmt.Println("这已经到达最后一页了")
		// } else {
		// 	fmt.Println(111)
		// 	// c.OnHTML(".main .list_page span", func(e *colly.HTMLElement) {
		// 	fmt.Println(222)
		// 	cNext := c.Clone()
		// 	nextURLs := e.ChildAttrs(".list_page a", "href")
		// 	newURL := "https://quanben.io" + nextURLs[len(nextURLs)-1]
		// 	fmt.Println("下一章 URL:", newURL)
		// 	GetContent(newURL, cNext)
		// 	fmt.Printf("-- Visiting next page %s\n", "https://quanben.io"+newURL)
		// 	cNext.Visit("https://quanben.io" + nextURL)
		// 	// })
		// 	return
		// }
		// fmt.Println("下一章 URL:", newURL)
		// fmt.Println("章节名为：", e.ChildText("h1"))
		return
	})
	c.Visit("https://quanben.io" + nextURL)
	// return
}

// type Novel struct {
// 	Name     string
// 	CoverURL string
// 	Summary  string
// 	Chapters []Chapter
// }

// type Chapter struct {
// 	Name    string
// 	Content string
// }

// func main() {
// 	baseUrl := "https://example.com" // 小说网站的首页地址
// 	novelUrls := make([]string, 0)   // 存储小说分类页面地址
// 	novels := make([]Novel, 0)       // 存储小说信息

// 	c := colly.NewCollector()

// 	// 访问小说首页，获取小说分类页面地址
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		href := e.Attr("href")
// 		if strings.Contains(href, "/category") {
// 			novelUrls = append(novelUrls, baseUrl+href)
// 		}
// 	})

// 	// 访问小说分类页面，获取小说列表页面地址
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		href := e.Attr("href")
// 		if strings.Contains(href, "/book") {
// 			go func() {
// 				novelUrl := baseUrl + href
// 				novelName := e.Text
// 				chapters := make([]Chapter, 0)
// 				summary := ""
// 				coverURL := ""

// 				// 访问小说详情页，获取小说信息
// 				detailCollector := colly.NewCollector()
// 				detailCollector.OnHTML("div.book-info", func(e *colly.HTMLElement) {
// 					summary = e.ChildText("p")
// 					coverURL = e.ChildAttr("img", "src")
// 				})
// 				detailCollector.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 					href := e.Attr("href")
// 					if strings.Contains(href, "/chapter") {
// 						go func() {
// 							chapterCollector := colly.NewCollector()
// 							chapterCollector.OnHTML("h1", func(e *colly.HTMLElement) {
// 								chapterName := e.Text
// 								chapterContent := ""
// 								chapterCollector.OnHTML("div.chapter-content", func(e *colly.HTMLElement) {
// 									chapterContent = e.Text
// 								})
// 								chapters = append(chapters, Chapter{Name: chapterName, Content: chapterContent})
// 							})
// 							chapterCollector.Visit(baseUrl + href)

// 						}()
// 					}
// 				})

// 				detailCollector.Visit(novelUrl)

// 				novels = append(novels, Novel{Name: novelName, CoverURL: coverURL, Summary: summary, Chapters: chapters})
// 			}()
// 		}
// 	})

// 	// 访问小说列表页面，获取小说详情页地址
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		href := e.Attr("href")
// 		if strings.Contains(href, "/book") {
// 			go func() {
// 				novelUrl := baseUrl + href
// 				novelName := e.Text
// 				chapters := make([]Chapter, 0)
// 				summary := ""
// 				coverURL := ""

// 				// 访问小说详情页，获取小说信息
// 				detailCollector := colly.NewCollector()
// 				detailCollector.OnHTML("div.book-info", func(e *colly.HTMLElement) {
// 					summary = e.ChildText("p")
// 					coverURL = e.ChildAttr("img", "src")
// 				})
// 				detailCollector.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 					href := e.Attr("href")
// 					if strings.Contains(href, "/chapter") {
// 						go func() {
// 							chapterCollector := colly.NewCollector()
// 							chapterCollector.OnHTML("h1", func(e *colly.HTMLElement) {
// 								chapterName := e.Text
// 								chapterContent := ""
// 								chapterCollector.OnHTML("div.chapter-content", func(e *colly.HTMLElement) {
// 									chapterContent = e.Text
// 								})
// 								chapters = append(chapters, Chapter{Name: chapterName, Content: chapterContent})
// 							})
// 							chapterCollector.Visit(baseUrl + href)

// 						}()
// 					}
// 				})

// 				detailCollector.Visit(novelUrl)

// 				novels = append(novels, Novel{Name: novelName, CoverURL: coverURL, Summary: summary, Chapters: chapters})
// 			}()
// 		}
// 	})

// 	// 判断小说列表是否有下一页，若有则访问下一页
// 	c.OnHTML("a.next-page", func(e *colly.HTMLElement) {
// 		if !e.HasClass("grey") {
// 			nextPage := e.Attr("href")
// 			c.Visit(baseUrl + nextPage)
// 		} else {
// 			fmt.Println("当前分类下的小说列表已经爬取")
// 		}
// 	})
// }
