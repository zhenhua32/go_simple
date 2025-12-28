package main

import "book2/ch2/confformat"

func main() {
	if err := confformat.MarshaALL(); err != nil {
		panic(err)
	}
	if err := confformat.UnmarshalAll(); err != nil {
		panic(err)
	}
	if err := confformat.OtherJSONExamples(); err != nil {
		panic(err)
	}
}
