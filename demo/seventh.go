package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest("get", "http://h5.imiaobige.com/novel/visit/272815/?d=json&r=0.6453096873674202", nil)
	if err != nil {
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Mobile Safari/537.36")
	req.Header.Add("Referer", "http://h5.imiaobige.com/read/272815/3968607.html")
	req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("Host", "h5.imiaobige.com")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
	regex := `<div>(?s:(.+?))</div>`
	c, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return
	}
	fmt.Println(c.FindAllStringSubmatch(string(bytes), -1))
}
