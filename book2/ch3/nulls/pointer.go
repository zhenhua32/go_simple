package nulls

import (
	"encoding/json"
	"fmt"
)

type ExamplePointer struct {
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

func PointerEncoding() error {
	e := ExamplePointer{}
	if err := json.Unmarshal([]byte(jsonbBlob), &e); err != nil {
		return err
	}
	fmt.Printf("pointer unmarshal, no age:%+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("pointer marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fullJsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("pointer unmarshal, with age=0,: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("pointer marshal, with age=0:", string(value))

	return nil
}
