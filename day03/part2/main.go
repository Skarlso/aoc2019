package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var dir = map[string][]int{
	"R": {0, 1},
	"L": {0, -1},
	"D": {-1, 0},
	"U": {1, 0},
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	w1 := lines[0]
	w2 := lines[1]
	// min := math.MaxInt64
	w1Split := strings.Split(w1, ",")
	// This is ugly but I'm tired, so I don't care. It got the job done.
	grid := make([][]int, 100000)
	for g := 0; g < 100000; g++ {
		grid[g] = make([]int, 100000)
	}
	startX := 50000
	startY := 50000
	currX := startX
	currY := startY
	w1Steps := 0
	// Fill in the first wire.
	for _, s := range w1Split {
		d, n := s[0], s[1:]
		distance, _ := strconv.Atoi(n)
		heading := dir[string(d)]
		for i := 0; i < distance; i++ {
			currY, currX = currY+heading[0], currX+heading[1]
			if grid[currY][currX] == 0 {
				grid[currY][currX] = w1Steps
			}
			w1Steps++
		}
	}

	// Fill in the second wire and record intersections and their distances.
	min := math.MaxInt64
	w2Split := strings.Split(w2, ",")
	currX = startX
	currY = startY
	w2Steps := 0
	for _, s := range w2Split {
		d, n := s[0], s[1:]
		distance, _ := strconv.Atoi(n)
		heading := dir[string(d)]
		for i := 0; i < distance; i++ {
			currY, currX = currY+heading[0], currX+heading[1]
			if grid[currY][currX] > 0 {
				m := grid[currY][currX] + w2Steps
				if m < min && m > 0 {
					min = m
				}
			}
			w2Steps++
		}
	}

	// +2 because I'm counting begin and end steps
	fmt.Println("min: ", min+2)
}
