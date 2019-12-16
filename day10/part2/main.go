package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
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
	mx := len(grid)
	my := len(grid[0])
	seen := make(map[float64][]point)
	as := make([]float64, 0)
	for _, l := range meteorLocations {
		if l.y != my || l.x != mx {
			a := ((math.Atan2(float64(l.y-my), float64(l.x-mx))) * (180 / math.Pi))
			// a := math.Floor(x*1000) / 1000
			a += 90
			if a < 0 {
				a += 360
			}

			seen[a] = append(seen[a], point{y: l.y, x: l.x})
			as = append(as, a)
		}
	}

	sort.Float64s(as)
	for _, s := range seen {
		sort.SliceStable(s, func(i, j int) bool {
			return (abs(s[i].y-my) + abs(s[i].x-mx)) < (abs(s[j].y-my) + abs(s[j].x-mx))
		})
	}
	count := 1
loop:
	for {
		for _, s := range as {
			var curr point
			curr, seen[s] = seen[s][0], seen[s][1:]
			if count == 200 {
				fmt.Println((curr.x * 100) + curr.y)
				break loop
			}
			count++
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
