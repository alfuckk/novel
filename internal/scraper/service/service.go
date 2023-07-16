package service

import (
	"fmt"
	"runtime"

	"github.com/gocolly/colly/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
)

type Service interface {
	Scrape(url string) (string, error)
}
type Params struct {
	conn *amqp.Connection
	fx.In
}
type service struct {
}

// 父函数
func (s service) Scrape(url string) (string, error) {
	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)
	// amqp.
	ch := amqp.Channel
	fmt.Println(ch)
	fmt.Println(c)
	tasks := make(chan string, 100)
	results := make(chan string, 100)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for item := range tasks {
				// 这里是处理任务的代码，例如访问URL并获取结果
				res := processURL(item)
				results <- res
			}
		}()
	}
	c.OnHTML(`div.nav a`, func(e *colly.HTMLElement) {
		tasks <- url + e.Attr("href")
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	close(tasks)

	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
	return "", nil
}

func NewService(p Params) Service {
	return service{
		amqp: p.conn,
	}
}

func processURL(url string) string {
	tasks := make(chan string, 100)
	results := make(chan string, 100)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			fmt.Println(url)
			for item := range tasks {
				// 这里是处理任务的代码，例如访问URL并获取结果
				res := subParseURL(item)
				results <- res
			}
		}()
	}

	c := colly.NewCollector()
	c.OnHTML(`.box .list2 a`, func(e *colly.HTMLElement) {
		// tasks <- url + e.Attr("href")
		tasks <- e.Attr("href")
		// fmt.Println(e.Text)
	})
	err := c.Visit(url)
	if err != nil {
		panic(err)
	}
	close(tasks)

	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}
	return url
}

func subParseURL(url string) (result string) {
	fmt.Println(url)
	return result
}
