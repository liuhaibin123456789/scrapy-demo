package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	err := c.Post("https://accounts.douban.com/passport/login?redir=https://www.douban.com/people/252833746/?_i=7518939OFLFag_", map[string]string{"username": "15922909035", "password": "13423964428"})
	if err != nil {
		fmt.Println("ssss")
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//获取标签的元素值
		attr := e.Attr("href")
		println(e.Name, ":", attr)
	})

	// start scraping
	err = c.Visit("https://www.douban.com/people/252833746/?_i=7518939OFLFag_")
	if err != nil {
		return
	}
}
