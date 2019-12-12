package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type point struct {
	x, y int
}

func main() {
	meteorLocations := make([]point, 0)
	file := os.Args[1]
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(content, []byte("\n"))
	grid := make([][]rune, 0)

	var meteors int
	for y, l := range lines {
		row := make([]rune, 0)
		for x, r := range l {
			row = append(row, rune(r))
			if r == '#' {
				meteors++
				meteorLocations = append(meteorLocations, point{y: y, x: x})
			}
		}
		grid = append(grid, row)
	}

	fmt.Println("Total number of meteors: ", meteors)
	directions := [][]int{
		// normal cases
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  //right
		// diagonal
		{-1, -1}, // upper-left
		{-1, 1},  // upper-right
		{1, -1},  // lower-left
		{1, 1},   // lower-right
		// the fringe cases
		// one step up two steps right
		{-1, -2}, // upper-left-1
		{-1, 2},  // upper-right-1
		{1, -2},  // lower-left-1
		{1, 2},   // lower-right-1
		// two steps up one step right
		{-2, -1}, // upper-left-2
		{-2, 1},  // upper-right-2
		{2, -1},  // lower-left-2
		{2, 1},   // lower-right-2
	}
	for _, loc := range meteorLocations {
		// temp meteor count starts at all the meteors
		tempM := meteors
		for _, d := range directions {
			tempLoc := loc
			var meteorSumInPath int
			for {
				tempLoc.y += d[0]
				tempLoc.x += d[1]

				// Make sure we are still in range
				if tempLoc.x < 0 || tempLoc.y < 0 || tempLoc.x >= len(grid[tempLoc.y]) || tempLoc.y >= len(grid) {
					break
				}

			}
		}
	}
}
