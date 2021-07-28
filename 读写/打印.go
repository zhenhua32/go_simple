package main

import (
	"fmt"
	"os"
)

type user struct {
	Name string
	Age  int
}

func print() {
	u := user{Name: "gopher", Age: 1}

	// 通用
	fmt.Println("通用占位符:")
	fmt.Printf("%v \n", u)
	fmt.Printf("%#v \n", u)
	fmt.Printf("%T \n", u)
	fmt.Printf("%% \n")
	// 布尔
	fmt.Println("布尔占位符:")
	fmt.Printf("%t \n", true)
	// 数字
	fmt.Println("数字占位符:")
	fmt.Printf("%b \n", 11)
	fmt.Printf("%c \n", 0x00004E2D)
	fmt.Printf("%d \n", 11)
	fmt.Printf("%o \n", 11)
	fmt.Printf("%q \n", 11)
	fmt.Printf("%x \n", 11)
	fmt.Printf("%X \n", 11)
	fmt.Printf("%U \n", 11)
	// 浮点数和复数
	fmt.Println("浮点数和复数占位符:")
	fmt.Printf("%b \n", 1.1)
	fmt.Printf("%e \n", 1.1)
	fmt.Printf("%E \n", 1.1)
	fmt.Printf("%f \n", 1.1)
	fmt.Printf("%F \n", 1.1)
	fmt.Printf("%g \n", 1.1)
	fmt.Printf("%G \n", 1.1)
	// 字符串
	fmt.Println("字符串占位符:")
	fmt.Printf("%s \n", "中国")
	fmt.Printf("%q \n", "中国")
	fmt.Printf("%x \n", "中国")
	fmt.Printf("%X \n", "中国")
	// 切片
	fmt.Println("切片占位符:")
	fmt.Printf("%p \n", []string{"a", "b"})
	// 指针
	fmt.Println("指针占位符:")
	fmt.Printf("%p \n", &[]string{"a", "b"})
}

func printString() {
	fmt.Printf("%10v \n", "hello")
	fmt.Printf("%010v \n", "hello")
	fmt.Printf("%-10v \n", "hello")
}

func printFunc() {
	u := user{Name: "gopher", Age: 1}
	fmt.Print(1, 2, "a", "b", "\n")
	fmt.Printf("name: %s, age: %d \n", u.Name, u.Age)
	fmt.Println("hello")
}

func printFuncF() {
	f, _ := os.Create("a.txt")
	u := user{Name: "gopher", Age: 1}
	fmt.Fprint(f, 1, 2, "a", "b", "\n")
	fmt.Fprintf(f, "name: %s, age: %d \n", u.Name, u.Age)
	fmt.Fprintln(f, "hello")
}

func printFuncS() {
	var a string
	u := user{Name: "gopher", Age: 1}
	a = fmt.Sprint(1, 2, "a", "b", "\n")
	fmt.Print(a)

	a = fmt.Sprintf("name: %s, age: %d \n", u.Name, u.Age)
	fmt.Print(a)

	a = fmt.Sprintln("hello")
	fmt.Print(a)
}

func scanFunc() {
	var a string
	n, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, a)

	var b int
	n, err = fmt.Scanf("%d", &b)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, b)

	var c string
	n, err = fmt.Scanln(&c)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, c)
}

func scanFuncF() {
	f, _ := os.Open("a.txt")
	var a string
	fmt.Fscan(f, &a)
	fmt.Println(a)

	var b string
	fmt.Fscanf(f, "%s", &b)
	fmt.Println(b)

	var c string
	fmt.Fscanln(f, &c)
	fmt.Println(c)
}

func scanFuncS() {
	var a string
	var b string
	fmt.Sscan("hello world", &a)
	fmt.Println(a)

	fmt.Sscanf("hello world", "%4s", &a)
	fmt.Println(a)

	fmt.Sscanln("hello \n world", &a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	scanFuncS()
}
