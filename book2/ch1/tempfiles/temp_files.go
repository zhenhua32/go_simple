package tempfiles

import (
	"fmt"
	"os"
)

func WorkWithTemp() error {
	t, err := os.MkdirTemp("", "tmp")
	if err != nil {
		return err
	}

	defer os.RemoveAll(t)
	fmt.Println("temp dir", t)

	tf, err := os.CreateTemp(t, "tmp")
	if err != nil {
		return err
	}
	fmt.Println(tf.Name())

	return nil

}
