package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

// 3 wide 2 tall -> 3 groups in 2 layers.
func main() {
	file := os.Args[1]
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	wide := 25
	tall := 6

	min := math.MaxInt64
	var minLayer string
	for i := 0; i < len(content); i += wide*tall {
		if i + (wide*tall) >= len(content) {
			break
		}
		layer :=  string(content[i:i+(wide*tall)])
		zeroCount := strings.Count(layer, "0")
		if zeroCount < min {
			minLayer = layer
			min = zeroCount
		}
	}
	fmt.Println(minLayer)
	ones := strings.Count(minLayer, "1")
	twos := strings.Count(minLayer, "2")
	fmt.Println("Mult: ", ones*twos)
}
