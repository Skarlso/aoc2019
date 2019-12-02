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
	bytesArr := bytes.Split(content, []byte(","))
	numbers := make([]int, len(bytesArr))
	for i := 0; i < len(bytesArr); i++ {
		n, _ := strconv.Atoi(string(bytesArr[i]))
		numbers[i] = n
	}
	// Initialising memory
	memory := make(map[int]int)
	inc := 4
	goal := 19690720
	var (
		noun, verb int
	)
	for one := 0; one < 100; one++ {
		for two := 0; two < 100; two++ {
			// Reset
			for i, v := range numbers {
				memory[i] = v
			}
			memory[1] = one
			memory[2] = two
			var i int
		loop:
			for {
				pos1, pos2, loc := memory[i+1], memory[i+2], memory[i+3]
				switch memory[i] {
				case 1:
					memory[loc] = memory[pos1] + memory[pos2]
				case 2:
					memory[loc] = memory[pos1] * memory[pos2]
				case 99:
					break loop
				}
				i += inc
			}
			if memory[0] == goal {
				noun = one
				verb = two
				fmt.Println("found it.")
				break
			}
		}
	}
	fmt.Println(100*noun + verb)
}
