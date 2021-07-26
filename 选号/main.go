package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func handlerError(err error) {
	if err != nil {
		panic(err)
	}
}

// 从网络获取 html
func getHtml() string {
	resp, err := http.Get("https://hangzhou.jihaoba.com/")

	handlerError(err)
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		handlerError(fmt.Errorf("status code is not 200, %v", resp.StatusCode))
	}

	html, err := io.ReadAll(resp.Body)
	handlerError(err)
	return string(html)
}

// 提取手机号
func extractPhone(html string) []string {
	// 创建正则表达式
	re, err := regexp.Compile(`(1\d{10})`)
	handlerError(err)

	// 查找所有匹配的字符串
	result := re.FindAllString(html, -1)
	return result
}

func main() {
	html := getHtml()

	phones := extractPhone(html)
	fmt.Println(len(phones))
	fmt.Println(phones)

}
