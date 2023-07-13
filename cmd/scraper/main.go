package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Novel struct {
	Title       string
	Author      string
	Description string
	Link        string
}

func main() {
	c := colly.NewCollector()

	c.OnHTML(".list .item", func(e *colly.HTMLElement) {
		novel := Novel{
			Title:       e.ChildText(".title"),
			Author:      e.ChildText(".author"),
			Description: e.ChildText(".desc"),
			Link:        e.Request.AbsoluteURL(e.ChildAttr("a", "href")),
		}
		fmt.Println(novel)
	})

	c.Visit("https://www.quanben.io/")
}
