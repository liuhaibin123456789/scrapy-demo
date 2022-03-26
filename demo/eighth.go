package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	//获取网站html
	rsp, err := http.Get("http://www.imiaobige.com/read/272815/3963108.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rsp.Body.Close()
	var p = make([]byte, 4089)
	var result string
	for true {
		n, err := rsp.Body.Read(p)
		if n == 0 {
			fmt.Println(n)
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		result += string(p[:n])
	}
	//fmt.Println(result)
	//提取数据
	regexStr := "<div id=\"content\">(?s:(.*?))</div>"
	c, err := regexp.Compile(regexStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	stringSubmatch := c.FindAllStringSubmatch(result, -1)
	fmt.Println(stringSubmatch[0][1])
}
