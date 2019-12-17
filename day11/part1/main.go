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

type robot struct {
	heading      []int
	headingIndex int
	location     point
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
	field := map[point]int{
		{0, 0}: 0,
	}
	headings := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	r := robot{
		heading:  []int{1, 0},
		location: point{0, 0},
	}
	m := intcode.NewMachine(memory)
	m.Input = []int{0}
	var (
		done bool
		out  []int
	)
	for !done {
		// Make the robot move.
		out, done = m.ProcessProgram()
		color, heading := out[0], out[1]
		if heading == 0 {
			r.headingIndex = (r.headingIndex + 1) % 4
		} else if heading == 1 {
			r.headingIndex = (r.headingIndex - 1)
			if r.headingIndex < 0 {
				r.headingIndex = 3
			}
		}
		r.heading = headings[r.headingIndex]
		field[r.location] = color
		r.move()
		m.Input = append(m.Input, field[r.location])
	}

	fmt.Println(len(field))
}

func (r *robot) move() {
	r.location.y += r.heading[0]
	r.location.x += r.heading[1]
}
