package main

import (
	"fmt"
	"regexp"

	"tzh.com/re/base"
)

func extractEmail(html string) []string {
	re, err := regexp.Compile(`\w+@[\w\.]+\.\w+`)
	base.HandlerError(err)

	result := re.FindAllString(html, -1)
	return result
}

func main() {
	html := base.GetHtml("https://www.sina.com.cn/contactus.html")

	emails := extractEmail(html)
	fmt.Println(len(emails))
	fmt.Println(emails)

	emails = base.Unique(emails)
	fmt.Println(len(emails))
	fmt.Println(emails)
}
