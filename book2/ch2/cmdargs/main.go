package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
	c := MenuConf{}
	menu := c.SetupMenu()

	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing params %s, error: %v", os.Args[1:], err)
		return
	}

	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubmenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				if err := f.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stdout, "Error parsing params %s, error: %v", os.Args[3:], err)
					return
				}
			}
			c.Greet(os.Args[2])
		default:
			fmt.Println("invalid command")
			menu.Usage()
			return
		}

	} else {
		menu.Usage()
		return
	}
}
