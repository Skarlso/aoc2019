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
	x, y     int
	distance int
	rot      int
	angle    float64
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
	mx := 19 // from previous part1 this is the location of the station
	my := 14
	seen := make(map[float64][]point)
	as := make([]point, 0)
	for _, l := range meteorLocations {
		if l.y != my || l.x != mx {
			a := ((math.Atan2(float64(l.y-my), float64(l.x-mx))) * (180 / math.Pi))
			a += 90
			if a < 0 {
				a += 360
			}

			d := abs(l.y-my) - abs(l.x-mx)
			p := point{y: l.y, x: l.x, angle: a, distance: d}
			seen[a] = append(seen[a], p)
			as = append(as, p)
		}
	}

	sort.SliceStable(as, func(i, j int) bool {
		if as[i].angle == as[j].angle {
			return as[i].distance < as[j].distance
		}
		return as[i].angle < as[j].angle
	})
	prev := -1.0
	rot := 0
	for i, r := range as {
		if r.angle == prev {
			rot++
		} else {
			rot = 0
		}

		as[i].rot = rot
		prev = r.angle
	}
	sort.Slice(as, func(i, j int) bool {
		if as[i].rot == as[j].rot {
			return as[i].angle < as[j].angle
		}
		return as[i].rot < as[j].rot
	})

	fmt.Printf("%+v", as[199].x*100+as[199].y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
