package base

import (
	"fmt"
	"io"
	"net/http"
)

func HandlerError(err error) {
	if err != nil {
		panic(err)
	}
}

// 从网络获取 html
func GetHtml(url string) string {
	resp, err := http.Get(url)

	HandlerError(err)
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		HandlerError(fmt.Errorf("status code is not 200, %v", resp.StatusCode))
	}

	html, err := io.ReadAll(resp.Body)
	HandlerError(err)
	return string(html)
}

// 去重
func Unique(arr []string) []string {
	unique := map[string]int{}
	for _, v := range arr {
		unique[v]++
	}
	result := []string{}
	for k := range unique {
		result = append(result, k)
	}
	return result
}
