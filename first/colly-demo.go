package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"time"
)

/*
	爬取网站小说的一个正文内容 地址直达 https://www.qb5.tw/book_45014/51463277.html
*/
func main() {
	//创建收集器
	c := colly.NewCollector()

	//限速
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob:   "www.qb5.tw/*",
		Delay:        time.Second * 10,
		Parallelism:  1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	////设置redis作为储存后端
	//redisStorge := redisstorage.Storage{
	//	Address:  "localhost:6379",
	//	Password: "",
	//	DB:       0,
	//}
	//err = c.SetStorage(&redisStorge)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer redisStorge.Client.Close()

	//随机设置browser user agent
	extensions.RandomUserAgent(c)

	////设置代理
	//proxySwitcher, err := proxy.RoundRobinProxySwitcher(
	//	"socks5://127.0.0.1:7890",
	//	"socks5://127.0.0.1:7899",
	//)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//c.SetProxyFunc(proxySwitcher)

	////加入队列
	//q, err := queue.New(1, &redisStorge)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting...")
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("status code: ", r.StatusCode, "\nerr: ", err)
	})

	//err = q.AddURL("https://www.qb5.tw/book_45014/51463277.html")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = q.Run(c)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	c.Visit("https://www.qb5.tw/book_45014/51463277.html")
}
