package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Skarlso/intcode"
)

const (
	n = iota + 1
	s
	w
	e
)

const (
	wall = iota
	moved
	oxygen
)

type point struct {
	x, y int
}

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

	//directions := map[int]point{
	//	n: {y: -1, x: 0},
	//	s: {y: 1, x: 0},
	//	w: {y: 0, x: -1},
	//	e: {y: 0, x: 1},
	//}

	// seen ? copy the machine with its current state and go from there.
	// I have to keep track of the already visited coordinates and somehow branch.
	// I need to be able to clone the machine.
	m := intcode.NewMachine(memory)
	m.Clone()
}
