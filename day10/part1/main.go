package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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
	//directions := [][]int{
	//	// normal cases
	//	{-1, 0}, // up
	//	{1, 0},  // down
	//	{0, -1}, // left
	//	{0, 1},  //right
	//	// diagonal
	//	{-1, -1}, // upper-left
	//	{-1, 1},  // upper-right
	//	{1, -1},  // lower-left
	//	{1, 1},   // lower-right
	//	// the fringe cases
	//	// one step up two steps right
	//	{-1, -2}, // upper-left-1
	//	{-1, 2},  // upper-right-1
	//	{1, -2},  // lower-left-1
	//	{1, 2},   // lower-right-1
	//	// two steps up one step right
	//	{-2, -1}, // upper-left-2
	//	{-2, 1},  // upper-right-2
	//	{2, -1},  // lower-left-2
	//	{2, 1},   // lower-right-2
	//}

	// Take the Bresenham's line algorithm and return all the points between two points.
	// Does two points should be between the meteor and the edge of the grid.
	// go through all those points and determine how many meteors there are
	// subtract, max.
	var max int
	var maxLoc point
	for _, loc := range meteorLocations {
		// temp meteor count starts at all the meteors
		tempM := meteors
		//fmt.Println("temp: ", tempM)
		for _, d := range directions {
			//fmt.Println("============== Going in direction: ", d)
			tempLoc := loc
			if loc.x == 5 && loc.y == 8 {
				time.Sleep(1 * time.Second)
				fmt.Println("Direction: ", d)
			}
			var meteorSumInPath int
			for {
				tempLoc.y += d[0]
				tempLoc.x += d[1]

				// Make sure we are still in range
				if tempLoc.x < 0 || tempLoc.y < 0 || tempLoc.y >= len(grid) || tempLoc.x >= len(grid[tempLoc.y]) {
					break
				}

				if loc.x == 5 && loc.y == 8 {
					fmt.Print(string(grid[tempLoc.y][tempLoc.x]))
				}
				if grid[tempLoc.y][tempLoc.x] == '#' {
					meteorSumInPath++
				}
			}
			if loc.x == 5 && loc.y == 8 {
				fmt.Println()
			}
			//fmt.Printf("At location %d, %d meteors in path: %d\n", loc.y, loc.x, meteorSumInPath)
			// subtract all the meteors found on the path -1, the first one.
			//fmt.Println(meteorSumInPath)
			if meteorSumInPath > 0 {
				meteorSumInPath--
			}
			if loc.x == 5 && loc.y == 8 {
				fmt.Println("Sum in path: ", meteorSumInPath)
			}
			tempM -= meteorSumInPath
			if loc.x == 5 && loc.y == 8 {
				fmt.Println("tempM", tempM)
			}
			//fmt.Println(tempM)
		}
		if tempM > max {
			max = tempM
			maxLoc = loc
		}
	}

	// -1 because self is not included
	fmt.Printf("%d max meteor(s) seen at location y:%d, x:%d\n", max, maxLoc.y, maxLoc.x)
}
