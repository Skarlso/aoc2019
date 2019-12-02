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
	i := 0

	// replacing as instructed:
	numbers[1] = 12
	numbers[2] = 2
loop:
	for {
		switch numbers[i] {
		case 1:
			pos1, pos2, loc := numbers[i+1], numbers[i+2], numbers[i+3]
			numbers[loc] = numbers[pos1] + numbers[pos2]
		case 2:
			pos1, pos2, loc := numbers[i+1], numbers[i+2], numbers[i+3]
			numbers[loc] = numbers[pos1] * numbers[pos2]

		case 99:
			break loop
		}
		i += 4
	}
	fmt.Println(numbers[0])
}
