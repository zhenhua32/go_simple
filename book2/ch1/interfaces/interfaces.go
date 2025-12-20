package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in io.ReadSeeker, out io.Writer) error {
	// w 是一个 MultiWriter，写入 w 的数据会同时写入 out (你的 buffer) 和 os.Stdout (屏幕)
	w := io.MultiWriter(out, os.Stdout)

	// 第一次写入：
	// 从 in 读取 "example"，写入 w。
	// 此时屏幕显示 "example"，buffer 中也有 "example"。
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	// 重置读取位置：
	// 将 in (ReadSeeker) 的光标移回开头 (0, 0)。
	in.Seek(0, 0)
	buf := make([]byte, 64)
	// 第二次写入：
	// 再次从 in 读取 "example"，写入 w。
	// 此时屏幕追加显示 "example" (变成 "exampleexample")，buffer 中也追加了 "example"。
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()

	return nil
}
