package main

import (
	"fmt"
	"strings"
	"unicode"
)

var content = `#########
#b.A.@.a#
#########`

var content2 = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`

type point struct {
	x, y int
}

var directions = []point{
	{y: -1, x: 0},
	{y: 1, x: 0},
	{y: 0, x: -1},
	{y: 0, x: 1},
}

func walk(current point, maze []string, keys map[byte]struct{}) []point {
	moves := make([]point, 0)
	for _, d := range directions {
		next := point{x: current.x + d.x, y: current.y + d.y}
		if maze[next.y][next.x] != '#' {
			if unicode.IsUpper(rune(maze[next.y][next.x])) {
				if _, ok := keys[byte(unicode.ToLower(rune(maze[next.y][next.x])))]; ok {
					moves = append(moves, next)
				}
			} else {
				moves = append(moves, next)
			}
		}
	}
	return moves
}

func main() {
	// content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(content, "\n")
	maze := make([]string, 0)
	for _, l := range lines {
		maze = append(maze, l)
	}

	var start point
	// Search starting point
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == '@' {
				start.x = j
				start.y = i
				break
			}
		}
	}
	fmt.Println("Starting location is: ", start)

	// Collect how many keys we have together
	allKeysCount := 0
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if unicode.IsLower(rune(maze[i][j])) {
				allKeysCount++
			}
		}
	}
	fmt.Printf("There are %d number of keys in this maze.", allKeysCount)

	// End goal is to visit all locations for now. But in fact, if we don't find any new keys, I guess.
	// We first, have to collect all keys and then keep track if we have them all.
	keys := make(map[byte]struct{})
	seen := make(map[point]struct{})
	queue := []point{start}
	var current point
	var steps int

	fmt.Println()
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		fmt.Print(string(maze[current.y][current.x]))
		if _, ok := seen[current]; !ok {
			seen[current] = struct{}{}
			if unicode.IsLower(rune(maze[current.y][current.x])) {
				keys[maze[current.y][current.x]] = struct{}{}
			}
			for _, next := range walk(current, maze, keys) {
				if _, ok := seen[next]; !ok {
					queue = append(queue, next)
					steps++
				}
			}
		}
	}
	fmt.Println()
	fmt.Println("Number of steps: ", steps)
}
