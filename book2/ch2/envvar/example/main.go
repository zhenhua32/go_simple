package main

import (
	"book2/ch2/envvar"
	"bytes"
	"fmt"
	"os"
)

type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	var err error
	tf, err := os.CreateTemp("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	secret := `{
		"secret": "so so secret"}`
	if _, err := tf.Write(bytes.NewBufferString(secret).Bytes()); err != nil {
		panic(err)
	}

	if err := os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err := os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err := envvar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains = ", secret)
	fmt.Println("EXAMPLE_VERSION = ", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE = ", os.Getenv("EXAMPLE_ISSAFE"))
	fmt.Printf("Final Config: %#v\n", c)
}
