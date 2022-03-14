package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
)

func main() {

}

// Used to record blog information
type blogInfo struct {
	title  string
	href   string
	author string
}

// Used to parse html
func parse(html string) {
	// Parse html
	root, _ := htmlquery.Parse(strings.NewReader(html))
	titleList := htmlquery.Find(root, `//*[@id="post_list"]/div/div[2]/h3/a/text()`)
	hrefList := htmlquery.Find(root, `//*[@id="post_list"]/div/div[2]/h3/a/@href`)
	authorList := htmlquery.Find(root, `//*[@id="post_list"]/div/div[2]/div/a/text()`)
	// Traverse the result
	for i := range titleList {
		blog := blogInfo{}
		blog.title = htmlquery.InnerText(titleList[i])
		blog.href = htmlquery.InnerText(hrefList[i])
		blog.author = htmlquery.InnerText(authorList[i])
		fmt.Println(blog)
	}
}
