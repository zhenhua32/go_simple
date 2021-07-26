package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

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
	fmt.Println(time.Now())
	time.Sleep(5 * time.Second)
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

func downloadImgAsync(url string, c chan int, a chan int) {
	go func() {
		// 会等待, 直到进去
		c <- 1
		downloadImg(url)
		<-c
		a <- 1
	}()
}

func main() {
	os.RemoveAll("./img")
	os.Mkdir("./img", 0755)

	html := base.GetHtml("https://ent.163.com/")
	urls := extractImg(html)

	c := make(chan int, 5) // 并发控制
	a := make(chan int, 3) // 整体控制

	fmt.Println(len(urls))
	for _, url := range urls {
		go downloadImgAsync(url, c, a)
	}

	for i := 0; i < len(urls); i++ {
		<-a
	}

	files, _ := os.ReadDir("./img")
	fmt.Println("下载的图片数量: ", len(files))
	if len(files) != len(urls) {
		fmt.Println("下载数量有缺失")
	}
}
