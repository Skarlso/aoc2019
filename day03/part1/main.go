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
	w1Split := strings.Split(w1, ",")
	grid := make(map[point]bool)
	start := point{0, 0}
	curr := start
	grid[start] = true
	for _, s := range w1Split {
		d, n := s[0], s[1:]
		distance, _ := strconv.Atoi(n)
		heading := dir[string(d)]
		for i := 0; i < distance; i++ {
			curr.y, curr.x = curr.y+heading[0], curr.x+heading[1]
			grid[curr] = true
		}
	}

	min := math.MaxInt64
	w2Split := strings.Split(w2, ",")
	curr = start
	for _, s := range w2Split {
		d, n := s[0], s[1:]
		distance, _ := strconv.Atoi(n)
		heading := dir[string(d)]
		for i := 0; i < distance; i++ {
			curr.y, curr.x = curr.y+heading[0], curr.x+heading[1]
			if grid[curr] {
				m := abs(start.x-curr.x) + abs(start.y-curr.y)
				if m < min && m > 1 {
					min = m
				}
			}
		}
	}

	fmt.Println("min: ", min)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
