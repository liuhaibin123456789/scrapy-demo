package main

import (
	"fmt"
	"regexp"
)

func main() {
	html := "<a>hhh</a><div>aaaaaa\n\rbbbbb</div></br><div>cccc\n\nddddd</div></br><div>eeeeee\n\nffff</div><div></div>"
	regex := `<div>(?s:(.*?))</div>`
	c, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(fmt.Sprint(err))
		return
	}
	fmt.Println(c.FindAllStringSubmatch(html, -1))
}
