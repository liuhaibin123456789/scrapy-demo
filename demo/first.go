package main

import (
	"fmt"
	"net/http"
)

//该例子用于简单的模拟了浏览器的请求，即爬虫原理
func main() {
	//创建一个http请求
	req, _ := http.NewRequest("GET", "https://www.baidu.com/", nil)
	//编写请求头
	req.Header.Add("UserAgent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
	//初始化一个客户端
	client := &http.Client{}
	//客户端发起请求
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	//获得回应
	fmt.Println(resp)
}
