package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"time"
)

/*
	爬取连接下某本小说：https://www.qb5.tw/book_45014/
*/
func main() {
	c := colly.NewCollector()
	//伪造客户端,随机用户代理
	extensions.RandomUserAgent(c)

	c1 := c.Clone()
	//限制规则
	err := c1.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "www.qb5.tw/*",
		Parallelism:  1,
		Delay:        10 * time.Second,
	})
	if err != nil {
		return
	}
	extensions.RandomUserAgent(c1)

	//异步处理
	c1.Async = true

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("status code: ", r.StatusCode, "\nerr: ", err)
	})

	c.OnHTML("dl.zjlist", func(e *colly.HTMLElement) {
		e.ForEach("dd", func(i int, e1 *colly.HTMLElement) {
			fmt.Println(e1.Text)
			href := e1.ChildAttr("a", "href")
			//放入上下文，用于两个收集器共享
			ctx := colly.NewContext()
			ctx.Put("title", e.Text)
			err := c1.Request("get", "https://www.qb5.tw/book_45014/"+href, nil, ctx, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
		})
	})
	c1.OnHTML("div#content", func(e *colly.HTMLElement) {
		fmt.Println(e.Request.Ctx.Get("title"))
		fmt.Println(e.Text)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("url:", r.Request.URL, "status code: ", r.StatusCode, "\nerr: ", err)
	})
	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("url:", r.Request.URL, "status code: ", r.StatusCode, "\nerr: ", err)
	})

	err = c.Visit("https://www.qb5.tw/book_45014/")
	if err != nil {
		fmt.Println(err)
		return
	}
	c1.Wait()
}
