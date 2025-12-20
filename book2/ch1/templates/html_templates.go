package templates

import (
	"fmt"
	"html/template"
	"os"
)

func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello, {{.Name}}</h1>")
	if err != nil {
		return err
	}

	err = t.Execute(os.Stdout, map[string]string{"Name": "skk"})
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))

	return nil
}
