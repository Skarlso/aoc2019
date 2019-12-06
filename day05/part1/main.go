package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	position = iota
	immediate
)

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	bytesArr := bytes.Split(content, []byte(","))
	memory := make(map[int]int)
	for i := 0; i < len(bytesArr); i++ {
		n, _ := strconv.Atoi(string(bytesArr[i]))
		memory[i] = n
	}
	inc := 4
	var (
		i int
	)
loop:
	for {
		opcode := memory[i]
		op, modes := getOpCodeAndModes(opcode)
		switch op {
		case 1:
			args := getArguments(3, i, modes, memory)
			memory[args[2]] = args[0] + args[1]
			inc = 4
		case 2:
			args := getArguments(3, i, modes, memory)
			memory[args[2]] = args[0] * args[1]
			inc = 4
		case 3:
			memory[memory[i+1]] = 1 // input is hardcoded as 1
			inc = 2
		case 4:
			output := memory[memory[i+1]]
			fmt.Println(output)
			inc = 2
		case 99:
			break loop
		}
		i += inc
	}
}

func getArguments(num, i int, modes []int, memory map[int]int) (args []int) {
	for p := 0; p < num; p++ {
		var m int
		if p >= len(modes) {
			m = 0
		} else {
			m = modes[p]
		}
		switch m {
		case position:
			// Because parameters that an instruction writes to is always in position mode.
			if p > 1 && p + 1 == num {
				args = append(args, memory[i+p+1])
			} else {
				args = append(args, memory[memory[i+p+1]])
			}
		case immediate:
			args = append(args, memory[i+p+1])
		}
	}
	return
}

func getOpCodeAndModes(opcode int) (o int, modes []int) {
	sop := strconv.Itoa(opcode)
	l := len(sop)
	if len(sop) == 1 {
		o, _ = strconv.Atoi(sop)
		return o, nil
	}
	o, _ = strconv.Atoi(sop[l-2:])
	smodes := sop[:l-2]
	for i := len(smodes) - 1; i >= 0; i-- {
		m, _ := strconv.Atoi(string(smodes[i]))
		modes = append(modes, m)
	}
	return
}
