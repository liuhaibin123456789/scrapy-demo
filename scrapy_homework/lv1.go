package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
	地址：http://xiaodiaodaya.cn/
*/

var latestUrl = make([]string, 0) //最新笑话的url
var baseDomain = "http://xiaodiaodaya.cn"

func main() {
	//网站顶级域名全称
	var html string
	c := http.Client{}
	//第一个请求获取首页信息，获取html
	req, err := http.NewRequest("GET", baseDomain, nil)
	if err != nil {
		fmt.Print(err)
		return
	}

	//伪装成浏览器
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Mobile Safari/537.36")
	response, err := c.Do(req)
	if err != nil {
		fmt.Print(err)
		return
	}

	//解析html，提取目标资源url
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	html = string(r)
	err = filter1(html)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(latestUrl)
	for _, url := range latestUrl {
		err := reqContentAndSave(url)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//表示获取最新笑话
func filter1(html string) error {
	//非贪婪子式匹配
	regex := `<a href="(?:(/article/.+?))">`
	c, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(err)
		return err
	}
	stringSubMatch := c.FindAllStringSubmatch(html, -1)
	//fmt.Println(stringSubMatch)
	for _, v := range stringSubMatch {
		//保存第二个子式的url
		latestUrl = append(latestUrl, v[1])
	}
	return nil
}

func reqContentAndSave(url string) error {
	//根据url请求html静态资源
	doMain := baseDomain + url
	var html string
	var title string
	var content string
	var result string
	fmt.Println(doMain)
	req, err := http.NewRequest("", doMain, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//伪装成浏览器
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Mobile Safari/537.36")

	c := http.Client{}
	response, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	////解析html，提取目标资源url
	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	html = string(r)
	fmt.Println(html)
	//非贪婪匹配标题文字
	regexTitle := `<h2 class="titleview">(?:(.+?))</h2>`
	tRex, err := regexp.Compile(regexTitle)
	if err != nil {
		fmt.Println(err)
		return err
	}
	matchTitle := tRex.FindAllStringSubmatch(html, -1)
	for _, v := range matchTitle {
		fmt.Println(v[1])
		title = v[1]
	}

	//非贪婪匹配内容文字
	regexContent := `</p>(?:(.+?))<span`
	cRex, err := regexp.Compile(regexContent)
	if err != nil {
		fmt.Println(err)
		return err
	}
	matchContent := cRex.FindAllStringSubmatch(html, -1)
	for _, v := range matchContent {
		content = v[1]
	}
	//去除br标签,合并标题和内容
	result = title + "\n\r" + strings.ReplaceAll(content, "<br>", "\n")
	result = title + "\n\r" + strings.ReplaceAll(result, "<br/>", "\n")
	//保存和标题名字一样的文件
	file, err := os.Create(title + ".txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(result))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
