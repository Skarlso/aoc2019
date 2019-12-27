package main

import (
	"bytes"
	"fmt"
	"github.com/Skarlso/intcode"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	n = iota + 1
	s
	w
	e
)

const (
	wall = iota
	moved
	oxygen
)

type point struct {
	x, y int
}

var seen = map[point]bool{
	{0, 0}: true,
}

var directions = map[int]point{
	n: {y: -1, x: 0},
	s: {y: 1, x: 0},
	w: {y: 0, x: -1},
	e: {y: 0, x: 1},
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

	//explore(m)
	// Start with North
	in := []int{n}
	m.Input = in
	// start by going up from position 0, 0
	found := explore(m, point{y: 0, x: 0})

	if found {
		fmt.Println("Oxygen found")
	}
}

func explore(m *intcode.Machine, currentPosition point) bool {
	var (
		found bool
	)
	for {
		clone := m.Clone()
		possibleMoves := make([]int, 0)
		for k, d := range directions {
			c := clone.Clone()
			c.Input = []int{k}
			out, done := c.ProcessProgram()
			if done {
				return out[0] == oxygen
			}
			if out[0] == oxygen {
				return true
			}
			p := point{y: currentPosition.y + d.y, x: currentPosition.x + d.x}
			if _, ok := seen[p]; !ok && out[0] != wall {
				seen[p] = true
				possibleMoves = append(possibleMoves, k)
			}
		}

		// if there is only one possible move, we move
		if len(possibleMoves) == 0 {
			// no more moves left
			return found
		} else if len(possibleMoves) == 1 {
			clone.Input = []int{possibleMoves[0]}
			// We don't care about the return because the clone
			// already tried it and if they would be an end
			// the clone would already have ended
			clone.ProcessProgram()
		} else if len(possibleMoves) > 1 {
			// We move in all directions
			for _, d := range possibleMoves {
				found = explore(m, point{y: currentPosition.y + directions[d].y, x: currentPosition.x + directions[d].x})
			}
			if found {
				return found
			}
		}
	}
	return found
}
