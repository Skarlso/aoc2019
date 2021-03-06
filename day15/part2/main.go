package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Skarlso/intcode"
)

const (
	n = iota + 1
	s
	w
	e
)

const (
	debug = false
)

const (
	wall = iota
	moved
	oxygen
)

type point struct {
	x, y int
}

var seen = map[point]int{
	{0, 0}: moved,
}

var directions = map[int]point{
	n: {y: -1, x: 0},
	s: {y: 1, x: 0},
	w: {y: 0, x: -1},
	e: {y: 0, x: 1},
}

func logDebug(args ...interface{}) {
	if debug {
		fmt.Println(args...)
	}
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	bytesArr := bytes.Split(content, []byte(","))
	memory := make(map[int]int)
	for i := 0; i < len(bytesArr); i++ {
		n, _ := strconv.Atoi(string(bytesArr[i]))
		memory[i] = n
	}

	// seen ? copy the machine with its current state and go from there.
	// I have to keep track of the already visited coordinates and somehow branch.
	// I need to be able to clone the machine.
	m := intcode.NewMachine(memory)

	// Start with North
	in := []int{n}
	m.Input = in
	// start by going up from position 0, 0
	explore(m, point{y: 0, x: 0})
	fmt.Println("Oxygen location at: ", oxygenLocation)

	max := 0
	for p, cell := range seen {
		if cell != moved {
			continue
		}
		start := p
		cameFrom := map[point]point{
			start: start,
		}
		path := []point{start}
		curr := point{}
		for len(path) > 0 {
			curr, path = path[0], path[1:]
			if curr == oxygenLocation {
				break
			}
			for _, next := range moves(curr) {
				if _, ok := cameFrom[next]; !ok {
					cameFrom[next] = curr
					path = append(path, next)
				}
			}
		}

		// Count the steps that it took to get there.
		count := 0
		current := oxygenLocation
		for current != start {
			current = cameFrom[current]
			count++
		}
		if count > max {
			max = count
		}
	}

	fmt.Println(max) // Furthest position in the maze.
}


func moves(p point) []point {
	ret := make([]point, 0)
	for _, d := range directions {
		next := point {x: p.x + d.x, y: p.y + d.y}
		if seen[next] != wall { // include oxygen
			ret = append(ret, next)
		}
	}
	return ret
}


var oxygenLocation point

func explore(m *intcode.Machine, currentPosition point) {
	clone := m.Clone()
	for {
		possibleMoves := make([]int, 0)
		for k, d := range directions {
			c := clone.Clone()
			c.Input = []int{k}
			out, done := c.ProcessProgram()
			c.InstructionCount++
			if done {
				logDebug("It finished running, and out is: ", out)
				return
			}
			p := point{y: currentPosition.y + d.y, x: currentPosition.x + d.x}
			if out[0] == oxygen {
				fmt.Println("Found the oxygen!!")
				oxygenLocation = p
				//seen[p] = oxygen
			}
			logDebug("Point is: ", p)
			logDebug("Out is: ", out)
			if _, ok := seen[p]; !ok && out[0] != wall {
				logDebug("Not a wall and have not seen it yet... Adding to moves.")
				seen[p] = out[0]
				possibleMoves = append(possibleMoves, k)
			} else if out[0] == wall {
				seen[p] = out[0]
			}
		}
		logDebug("Possible moves: ", possibleMoves)

		// if there is only one possible move, we move
		if len(possibleMoves) == 0 {
			// no more moves left
			logDebug("No more moves left...")
			return
		} else if len(possibleMoves) == 1 {
			d := possibleMoves[0]
			clone.Input = []int{d}
			// We don't care about the return because the clone
			// already tried it and if they would be an end
			// the clone would already have ended
			clone.InstructionCount++
			clone.ProcessProgram()
			// Update the location we moved to.
			currentPosition.y += directions[d].y
			currentPosition.x += directions[d].x
		} else if len(possibleMoves) > 1 {
			//count := 0
			// We move in all directions
			for _, d := range possibleMoves {
				c := clone.Clone()
				c.Input = []int{d}
				c.ProcessProgram()
				p := point{y: currentPosition.y + directions[d].y, x: currentPosition.x + directions[d].x}
				explore(&c, p)
			}
		}
	}
}
