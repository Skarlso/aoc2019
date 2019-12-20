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
	memory[0] = 2

	grid := [24][42]int{}
	m := intcode.NewMachine(memory)

	var (
		done      bool
		out       []int
		segment   []int
		ballLoc   int
		paddleLoc int
	)
	out, done = m.ProcessProgram() // setup

	for len(out) > 0 {
		segment, out = out[:3], out[3:]
		if segment[0] == -1 {
			continue
		}
		grid[segment[1]][segment[0]] = segment[2]
		if segment[2] == paddle {
			paddleLoc = segment[1]
		}
		if segment[2] == ball {
			ballLoc = segment[1]
		}
	}
	drawGrid(grid)
	m.Input = []int{0}
	for {
		out, done = m.ProcessProgram()
		if done {
			fmt.Println(out) // out contains the final score.
			break
		}

		for len(out) > 0 {
			segment, out = out[:3], out[3:]
			switch segment[2] {
			case empty:
				grid[segment[1]][segment[0]] = empty
			case wall:
				grid[segment[1]][segment[0]] = wall
			case block:
				grid[segment[1]][segment[0]] = block
			case paddle:
				grid[segment[1]][segment[0]] = paddle
				paddleLoc = segment[0]
			case ball:
				grid[segment[1]][segment[0]] = ball
				ballLoc = segment[0]
				if ballLoc < paddleLoc {
					// fmt.Println("-1")
					m.Input = append(m.Input, -1)
				} else if ballLoc > paddleLoc {
					// fmt.Println("1")
					m.Input = append(m.Input, 1)
				} else if ballLoc == paddleLoc {
					// fmt.Println("0")
					m.Input = append(m.Input, 0)
				}
			}
		}
		drawGrid(grid)
	}
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
