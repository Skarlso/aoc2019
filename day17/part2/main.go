package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Skarlso/intcode"
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
	memory[0] = 2 // Wake up the robot
	m := intcode.NewMachine(memory)
	m.Input = []int{0}
	var (
		out  []int
		done bool
		// grid = make(map[point]int)
		// x, y int
	)

	for !done {
		out, done = m.ProcessProgram()
		display(out)
	}
}

func display(out []int) {
	for _, v := range out {
		fmt.Print(string(v))
	}
}
