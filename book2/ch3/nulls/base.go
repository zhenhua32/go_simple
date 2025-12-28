package nulls

import (
	"encoding/json"
	"fmt"
)

const (
	jsonbBlob    = `{"name": "Aaron"}`
	fullJsonBlob = `{"name": "Aaron", "age": 0}`
)

type Example struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func BaseEncoding() error {
	e := Example{}
	if err := json.Unmarshal([]byte(jsonbBlob), &e); err != nil {
		return err
	}
	fmt.Printf("regular unmarshal, no age:%+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("regular marshal, no age:", string(value))

	if err := json.Unmarshal([]byte(fullJsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("regular unmarshal, with age=0:%+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("regular marshal, with age=0:", string(value))

	return nil
}
