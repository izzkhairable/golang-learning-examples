package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	// Fill the slice with 'A' -> ASCII value 65
	for i := range b {
		b[i] = 'A'
	}
	// Return slice length and nil error
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
