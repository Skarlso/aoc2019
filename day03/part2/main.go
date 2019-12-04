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

type point struct {
	x, y int
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(content), "\n")
	w1 := lines[0]
	w2 := lines[1]
	// min := math.MaxInt64
	w1Split := strings.Split(w1, ",")
	grid := make(map[point]int)
	start := point{0, 0}
	curr := start
	w1Steps := 0
	// Fill in the first wire.
	for _, s := range w1Split {
		d, n := s[0], s[1:]
		distance, _ := strconv.Atoi(n)
		heading := dir[string(d)]
		for i := 0; i < distance; i++ {
			curr.y, curr.x = curr.y+heading[0], curr.x+heading[1]
			if grid[curr] == 0 {
				grid[curr] = w1Steps
			}
			w1Steps++
		}
	}

	// Fill in the second wire and record intersections and their distances.
	min := math.MaxInt64
	w2Split := strings.Split(w2, ",")
	curr = start
	w2Steps := 0
	for _, s := range w2Split {
		d, n := s[0], s[1:]
		distance, _ := strconv.Atoi(n)
		heading := dir[string(d)]
		for i := 0; i < distance; i++ {
			curr.y, curr.x = curr.y+heading[0], curr.x+heading[1]
			if grid[curr] > 0 {
				m := grid[curr] + w2Steps
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
