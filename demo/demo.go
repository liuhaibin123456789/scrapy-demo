package main

import (
	"fmt"
	"regexp"
)

func main() {
	html := "\t\t<div id=\"content\">ssssssssss你哈sssssssssss\n你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈你哈ssssssssssssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈ssssssssss你哈你哈ssssssssss你哈ssssssssss你哈sssssssss你哈ssssssssss你哈</div>"

	regex := `<div id="content">(?s:(.*?))</div>`
	c, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return
	}
	fmt.Println(c.FindAllStringSubmatch(html, -1))
}