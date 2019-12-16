package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

	max := 0
	for _, l := range meteorLocations {
		seen := make(map[float64]bool)

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == '#' && (y != l.y || x != l.x) {
					a := (math.Atan2(float64(y-l.y), float64(x-l.x))) * (180 / math.Pi)
					seen[a] = true
				}
			}
		}
		if len(seen) > max {
			max = len(seen)
		}
	}
	fmt.Println("max: ", max)
}
