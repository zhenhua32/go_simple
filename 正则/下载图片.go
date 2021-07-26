package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"tzh.com/re/base"
)

func extractImg(html string) []string {
	img_urls := []string{}
	// 核心是 ([^'"\s]+?\.(?:jpg|jpeg|png|gif))
	reg := `['"](?:http\:|https\:)?(?://)?([^'"\s]+?\.(?:jpg|jpeg|png|gif)[^'"\s]*?)['"]`

	re := regexp.MustCompile(reg)
	matches := re.FindAllStringSubmatch(html, -1)
	for _, match := range matches {
		img_urls = append(img_urls, match[1])
	}
	return img_urls
}

func downloadImg(url string) {
	names := strings.Split(url, "/")
	filename := "./img/" + strings.Split(strings.Split(names[len(names)-1], "&")[0], "?")[0]

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("下载失败: %v, error: %v \n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("下载失败: %v, 状态码是 %v", url, resp.StatusCode)
		return
	}

	img, err := io.ReadAll(resp.Body)
	base.HandlerError(err)

	err = os.WriteFile(filename, img, 0644)
	if err != nil {
		fmt.Printf("下载失败: %v, error: %v \n", url, err)
		return
	}
}

func main() {
	html := base.GetHtml("https://ent.163.com/")
	urls := extractImg(html)

	fmt.Println(len(urls))
	for _, url := range urls {
		downloadImg(url)
	}

	files, _ := os.ReadDir("./img")
	if len(files) != len(urls) {
		fmt.Println("下载数量有缺失")
	}
}
