package base

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

// 下载并保存图片
func DownloadImg(url string, filename string) {
	resp, err := http.Get(url)
	HandlerError(err)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		HandlerError(fmt.Errorf("status code is not 200, %v", resp.StatusCode))
		return
	}

	img, err := io.ReadAll(resp.Body)
	HandlerError(err)

	err = os.WriteFile(filename, img, 0644)
	if err != nil {
		HandlerError(fmt.Errorf("保存文件失败: %v, error: %v", url, err))
		return
	}
	fmt.Printf("下载成功: %v, 保存地址: %v \n", url, filename)
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
