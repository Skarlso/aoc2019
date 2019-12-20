package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Skarlso/intcode"
)

const (
	empty = iota
	wall
	block
	paddle
	ball
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	bytesArr := bytes.Split(content, []byte(","))

	memory := make(map[int]int)
	for i := 0; i < len(bytesArr); i++ {
		n, _ := strconv.Atoi(string(bytesArr[i]))
		memory[i] = n
	}

	m := intcode.NewMachine(memory)
	out, _ := m.ProcessProgram()

	var segment []int
	blockNum := 0
	for len(out) > 0 {
		segment, out = out[:3], out[3:]
		if segment[2] == block {
			blockNum++
		}
	}
	fmt.Println(blockNum)
}
