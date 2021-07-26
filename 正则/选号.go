package main

import (
	"fmt"
	"regexp"

	"tzh.com/re/base"
)

// 提取手机号
func extractPhone(html string) []string {
	// 创建正则表达式
	re, err := regexp.Compile(`(1\d{10})`)
	base.HandlerError(err)

	// 查找所有匹配的字符串
	result := re.FindAllString(html, -1)
	return result
}

func main() {
	html := base.GetHtml("https://hangzhou.jihaoba.com/")

	phones := extractPhone(html)
	fmt.Println(len(phones))
	fmt.Println(phones)

}
