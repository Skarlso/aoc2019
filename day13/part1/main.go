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

	// I looked and the largest number was 41x23
	// height := 41
	// width := 23
	grid := [24][42]int{}

	var segment []int
	blockNum := 0
	for len(out) > 0 {
		segment, out = out[:3], out[3:]
		// if a ball meets a block, it needs to break said block
		// fmt.Println(segment)
		switch segment[2] {
		case empty:
			grid[segment[1]][segment[0]] = empty
		case wall:
			grid[segment[1]][segment[0]] = wall
		case block:
			grid[segment[1]][segment[0]] = block
			blockNum++
		case paddle:
			grid[segment[1]][segment[0]] = paddle
		case ball:
			if grid[segment[1]][segment[0]] == block {
				grid[segment[1]][segment[0]] = empty
			}
			grid[segment[1]][segment[0]] = ball
		}
		drawGrid(grid)
	}
	drawGrid(grid)

	fmt.Println(blockNum)
}

func drawGrid(grid [24][42]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case empty:
				fmt.Print(" ")
			case wall:
				fmt.Print("#")
			case block:
				fmt.Print("▀")
			case paddle:
				fmt.Print("-")
			case ball:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}
