package main

import (
	"fmt"
	"regexp"

	"tzh.com/re/base"
)

func extractImg(html string) []string {
	img_urls := []string{}
	// 核心是 ([^'"\s]+?\.(?:jpg|jpeg|png|gif))
	reg := `['"](?:http\:|https\:)?(?://)?([^'"\s]+?\.(?:jpg|jpeg|png|gif))[^'"\s]*?['"]`

	re := regexp.MustCompile(reg)
	matches := re.FindAllStringSubmatch(html, -1)
	for _, match := range matches {
		img_urls = append(img_urls, match[1])
	}
	return img_urls
}

func main() {
	html := base.GetHtml("https://ent.163.com/")

	urls := extractImg(html)

	fmt.Println(len(urls))
	for _, url := range urls {
		fmt.Println(url)
	}
}
