// all programs start running in package main
package main

import (
    //"fmt" //'formatted IO'
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {

    var fields []string = strings.Fields(s)
    var words = make(map[string]int)

    for i := 0; i < len(fields); i++ {
        words[fields[i]]++
    }

    return words
}

func main() {
    wc.Test(WordCount)
}
