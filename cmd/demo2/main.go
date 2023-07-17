package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/gocolly/colly/v2"
)

type Category struct {
	Title string
	URL   string
	ID    int
}

type Novel struct {
	CategoryID    int
	Title         string
	Author        string
	Cover         string
	Status        string
	Desc          string
	LatestChapter int
}

type Chapter struct {
	NovelID int
	Title   string
	URL     string
	Content string
}

var API_URL = "https://quanben.io"

func main() {
	c := colly.NewCollector()
	var wg sync.WaitGroup
	category := make(chan Category)
	// results := make(chan Result)
	go ReadNavCategory(&category, c)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(2)
		go ReadNovelList(&category, &wg)
	}

	wg.Wait()

	fmt.Println("All workers done!")
}

func GetCategory() {

}

func ReadNavCategory(categorys *chan Category, cIndex *colly.Collector) {
	cIndex.OnHTML(".nav", func(e *colly.HTMLElement) {
		urls := e.ChildAttrs("a", "href")
		title := e.ChildTexts("a span")
		id := 1
		for i := 0; i <= len(urls)-1; i++ {
			item := Category{
				URL:   API_URL + urls[i],
				Title: title[i],
				ID:    id,
			}
			id++
			*categorys <- item
		}
	})

	cIndex.Visit(API_URL)
}

func ReadNovelList(categorys *chan Category, wg *sync.WaitGroup) {
	defer wg.Done()
	c := colly.NewCollector()
	if categorys != nil {
		for v := range *categorys {
			// cNovel := c.Clone()
			//	爬到此处，需要将分类信息存入数据库
			fmt.Printf("->Category:%d, Title:%s, URL:%s\n", v.ID, v.Title, v.URL)
			c.OnHTML(".nav", func(e *colly.HTMLElement) {
				urls := e.ChildAttrs("a", "href")
				title := e.ChildTexts("a span")
				fmt.Println(urls, title)
				c.Visit(v.URL)
			})
		}
	}
	wg.Wait()

}
