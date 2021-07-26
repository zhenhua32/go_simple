package main

import (
	"fmt"
	"regexp"

	"tzh.com/re/base"
)

func extractLink(html string) []string {
	// .*? 是非贪婪模式的 .*
	re, err := regexp.Compile(`<a href=["'](http.*?)["']`)
	base.HandlerError(err)
	res := re.FindAllStringSubmatch(html, -1)

	result := []string{}
	for _, v := range res {
		// 第 0 项匹配是整个匹配的字符串，第 1 项匹配是第一个子串, 即 (.*?) 里的内容
		result = append(result, v[1])
	}
	return result
}

func main() {
	html := base.GetHtml("https://touduyu.com/")

	links := extractLink(html)
	for _, v := range links {
		fmt.Println(v)
	}
}
