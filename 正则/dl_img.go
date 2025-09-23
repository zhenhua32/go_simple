package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
	"tzh.com/re/base"
)

// 获取所有初始链接 https://www.tuiimg.com/meinv/list_1.html
func getImageList(url string, end int) []string {
	initList := []string{}
	// 最多 155 页
	for i := 1; i <= end; i++ {
		initList = append(initList, url+fmt.Sprintf("list_%v.html", i))
	}

	return initList
}

// 解析图片列表页面
func extractDetailList(url string) []string {
	html := base.GetHtml(url)

	// 定义要解析的正则
	reg := `https://www.tuiimg.com/meinv/\d+/`
	re := regexp.MustCompile(reg)
	matches := re.FindAllStringSubmatch(html, -1)

	result := []string{}
	for _, match := range matches {
		result = append(result, match[0])
	}
	// 去重
	result = base.Unique(result)
	// 升序
	slices.Sort(result)
	return result
}

func main() {
	os.MkdirAll("./img", 0755)

	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}

	// url := "https://i.tuiimg.net/007/3084/4.jpg"
	// base.DownloadImg(url, "./img/4.jpg")

	initList := getImageList("https://www.tuiimg.com/meinv/", 10)
	fmt.Println(strings.Join(initList, "\n"))

	detailList := extractDetailList("https://www.tuiimg.com/meinv/list_1.html")
	fmt.Println(strings.Join(detailList, "\n"))
	fmt.Println(len(detailList))

}
