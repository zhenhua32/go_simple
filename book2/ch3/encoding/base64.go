package encoding

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
)

func Base64Example() error {
	value := base64.URLEncoding.EncodeToString([]byte("encoding some data!"))
	fmt.Println("With EncodingToString and URLEncoding:", value)
	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With DecodeToString and URLEncoding:", string(decoded))
	return nil
}

func Base64ExampleEncoder() error {
	buffer := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)

	if _, err := encoder.Write([]byte("encoding some other data")); err != nil {
		return err
	}
	if err := encoder.Close(); err != nil {
		return err
	}
	fmt.Println("Using encoder and StdEncoding", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	results, err := io.ReadAll(decoder)
	if err != nil {
		return err
	}
	fmt.Println("Using decoder and StdEncoding:", string(results))
	return nil
}
