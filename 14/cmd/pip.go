package main

import (
	"io"
)

func write(w io.Writer) error {
	metadata := []byte{1, 2, 3}
	_, err := w.Write(metadata)
	if err != nil {
		return err
	}

	_, err = w.Write(metadata)
	if err != nil {
		return err
	}

	_, err = w.Write(metadata)
	if err != nil {
		return err
	}

	_, err = w.Write(metadata)
	if err != nil {
		return err
	}

	return nil
}
