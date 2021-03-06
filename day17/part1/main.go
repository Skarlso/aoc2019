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

	m := intcode.NewMachine(memory)
	m.Input = []int{0}
	out, _ := m.ProcessProgram()
	// if done {
	// 	display(out)
	// }
	var (
		grid = make(map[point]int)
		x, y int
	)
	for _, v := range out {
		grid[point{x: x, y: y}] = v
		switch v {
		case 10:
			y++
			x = 0
		case 35:
			x++
		case 46:
			x++
		}
	}
	sum := 0
	for k, v := range grid {
		if v == 35 {
			// check up, right, down, left
			if grid[point{x: k.x, y: k.y - 1}] == 35 && grid[point{x: k.x + 1, y: k.y}] == 35 && grid[point{x: k.x, y: k.y + 1}] == 35 && grid[point{x: k.x - 1, y: k.y}] == 35 {
				sum += k.x * k.y
				grid[k] = 'O'
			}
		}
	}
	fmt.Println(sum)
}

func display(out []int) {
	for _, v := range out {
		fmt.Print(string(v))
	}
}
