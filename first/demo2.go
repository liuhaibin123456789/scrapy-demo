package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
	"os"
	"time"
)

/*
	爬取连接下某本小说：https://www.qb5.tw/book_45014/
*/
func main() {
	c := colly.NewCollector()
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "www.qb5.tw/*",
		Parallelism:  1,
		Delay:        5 * time.Second,
	})
	//伪造客户端,随机用户代理
	extensions.RandomUserAgent(c)

	c1 := c.Clone()
	//限制规则
	err = c1.Limit(&colly.LimitRule{
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

	//创建文件对象
	file, err := os.Create("万象之王.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("status code: ", r.StatusCode, "\nerr: ", err)
	})

	c.OnHTML("dl.zjlist", func(e *colly.HTMLElement) {
		e.ForEach("dd", func(i int, e1 *colly.HTMLElement) {
			//获取指定标签的属性值
			href := e1.ChildAttr("a", "href")
			c1.Visit("https://www.qb5.tw/book_45014/" + href)
			//err := c1.Request("get", "https://www.qb5.tw/book_45014/"+href, nil, nil, nil)
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}
		})
	})
	c1.OnHTML("div#content", func(e *colly.HTMLElement) {
		//通过相邻标签元素进行定位
		title := e.ChildText("div.nav-style+h1")
		n, err := file.Write([]byte("\t\t\t\t\t\t" + title + "\n\r" + e.Text + "\n\r"))
		if err != nil {
			fmt.Println(n, err)
			return
		}
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("c:", "url:", r.Request.URL, "status code: ", r.StatusCode, "\nerr: ", err)
	})
	c1.OnError(func(r *colly.Response, err error) {
		fmt.Println("c1:", "url:", string(r.Body), r.Request.URL, "status code: ", r.StatusCode, "\nerr: ", err)
	})
	//内存队列
	q, _ := queue.New(5, &queue.InMemoryQueueStorage{MaxSize: 10000})
	err = q.AddURL("https://www.qb5.tw/book_45014/")
	if err != nil {
		return
	}
	err = q.Run(c)
	if err != nil {
		return
	}
	c1.Wait()
}
