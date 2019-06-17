// all programs start running in package main
package main

import (
    "golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader
func (mr MyReader) Read(b []byte) (int, error) {
    for i := 0; i < 8; i++ {
        b[i] = 'A'
    }

    return 8, nil
}

func main() {
    reader.Validate(MyReader{})
}
