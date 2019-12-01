package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := bytes.Split(content, []byte("\n"))
	var sum int
	for _, l := range lines {
		i, _ := strconv.Atoi(string(l))
		val := (i / 3) - 2
		for v := (val / 3) - 2; v > 0; v = (v / 3) - 2 {
			val += v
		}
		sum += val
	}
	fmt.Println(sum)
}
