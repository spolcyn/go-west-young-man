// all programs start running in package main
package main

import (
    "io"
    "os"
    "strings"
    //"fmt"
)

type rot13Reader struct {
    r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
    rot13.r.Read(b);

    for i := 0; i < len(b); i++ {
        if b[i] >= 65 && b[i] <= 90 {
            b[i] += 13
            if b[i] > 90 {
                b[i] -= 26
            }
        } else if b[i] >= 97 && b[i] <= 122 {
            b[i] += 13
            if b[i] > 122 {
                b[i] -= 26
            }
        }
    }

    return len(b), nil
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
