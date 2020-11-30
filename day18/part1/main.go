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

type bot struct {
	loc   point
	keys  map[byte]struct{}
	steps int
}

var directions = []point{
	{y: -1, x: 0},
	{y: 1, x: 0},
	{y: 0, x: -1},
	{y: 0, x: 1},
}

func (b *bot) walk(maze []string) []bot {
	moves := make([]point, 0)
	for _, d := range directions {
		next := point{x: b.loc.x + d.x, y: b.loc.y + d.y}
		if maze[next.y][next.x] != '#' {
			if unicode.IsUpper(rune(maze[next.y][next.x])) {
				if _, ok := b.keys[byte(unicode.ToLower(rune(maze[next.y][next.x])))]; ok {
					moves = append(moves, next)
				}
			} else if unicode.IsLower(rune(maze[next.y][next.x])) {
				b.keys[maze[next.y][next.x]] = struct{}{}
				moves = append(moves, next)
			} else {
				moves = append(moves, next)
			}
		}
	}
	bots := make([]bot, 0)
	for _, m := range moves {
		newBot := b.clone()
		newBot.loc = m
		bots = append(bots, newBot)
	}
	return bots
}

func (b *bot) clone() bot {
	a := &bot{}
	a.keys = make(map[byte]struct{})
	for k, v := range b.keys {
		a.keys[k] = v
	}
	a.loc = b.loc
	a.steps = b.steps
	return *a
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

	starterBot := bot{loc: start, keys: make(map[byte]struct{})}
	// End goal is to visit all locations for now. But in fact, if we don't find any new keys, I guess.
	// We first, have to collect all keys and then keep track if we have them all.
	queue := []bot{starterBot}
	var current bot
	fmt.Println()
	// Keep track of the bots? But how?
	// endless loop until all keys are found
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if len(current.keys) == allKeysCount {
			fmt.Printf("The first bot which found all the keys took %d steps.\n", current.steps)
			break
		}
		queue = append(queue, current.walk(maze)...)
	}
	fmt.Println()
}
