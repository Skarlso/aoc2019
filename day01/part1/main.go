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
	var sum int64
	for _, l := range lines {
		i, _ := strconv.Atoi(string(l))
		sum += int64((i / 3) - 2)
	}
	fmt.Println(sum)
}
