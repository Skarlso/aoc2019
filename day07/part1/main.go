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
	position = iota
	immediate
)

func main() {
	filename := os.Args[1]
	//input := os.Args[2]
	//in, _ := strconv.Atoi(input)
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

	sequence := []int{0, 1, 2, 3, 4}
	sequences := permutation(sequence)

	machineA := intcode.NewMachine()
	machineB := intcode.NewMachine()
	machineC := intcode.NewMachine()
	machineD := intcode.NewMachine()
	machineE := intcode.NewMachine()

	var max int
	for _, seq := range sequences {
		m := memClone(memory)
		machineA.Memory = m
		machineA.Input = []int{seq[0], 0}
		machineA.ProcessProgram()
		out := processProgram("A", []int{seq[0], 0}, m)
		m = memClone(memory)
		machineB.Memory = m
		machineB.Input = []int{seq[1], machineA.Output[0]}
		out = processProgram("B", []int{seq[1], out}, m)
		m = memClone(memory)
		machineC.Memory = m
		machineC.Input = []int{seq[2], machineB.Output[0]}
		out = processProgram("C", []int{seq[2], out}, m)
		m = memClone(memory)
		machineD.Memory = m
		machineD.Input = []int{seq[3], machineC.Output[0]}
		out = processProgram("D", []int{seq[3], out}, m)
		m = memClone(memory)
		machineE.Memory = m
		machineE.Input = []int{seq[4], machineD.Output[0]}
		if out > max {
			max = out
		}
	}
	fmt.Println("Max output: ", max)
}

func memClone(memory map[int]int) map[int]int {
	m := make(map[int]int)
	for k, v := range memory {
		m[k] = v
	}
	return m
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)
	return
}

func processProgram(amp string, phases []int, memory map[int]int) (out int) {
	var (
		i int
	)
	inputIndex := 0
loop:
	for {
		inc := 4
		opcode := memory[i]
		//fmt.Println("opcode: ", opcode)
		op, modes := getOpCodeAndModes(opcode)
		//fmt.Println(memory)
		//time.Sleep(1 * time.Second)
		//fmt.Println("i, op: ", i, op)
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
			memory[memory[i+1]] = phases[inputIndex] // input is hardcoded as 5
			inc = 2
			inputIndex++
		case 4:
			var output int
			if len(modes) > 0 {
				switch modes[0] {
				case position:
					output = memory[memory[i+1]]
				case immediate:
					output = memory[i+1]
				}
			} else {
				output = memory[memory[i+1]]
			}
			out = output
			inc = 2
		case 5:
			args := getArguments(2, i, modes, memory)
			if args[0] != 0 {
				i = args[1]
				inc = 0
			} else {
				inc = 3
			}
			//fmt.Printf("5 i: %d args: %+v\n", i, args)
		case 6:
			args := getArguments(2, i, modes, memory)
			if args[0] == 0 {
				i = args[1]
				inc = 0
			} else {
				inc = 3
			}
			//fmt.Printf("6 i: %d args: %+v\n", i, args)
		case 7:
			args := getArguments(3, i, modes, memory)
			if args[0] < args[1] {
				memory[args[2]] = 1
			} else {
				memory[args[2]] = 0
			}
			//fmt.Printf("7 i: %d args: %+v\n", i, args)
			inc = 4
		case 8:
			args := getArguments(3, i, modes, memory)
			if args[0] == args[1] {
				memory[args[2]] = 1
			} else {
				memory[args[2]] = 0
			}
			//fmt.Printf("8 i: %d args: %+v\n", i, args)
			inc = 4
		case 99:
			break loop
		}
		i += inc
	}

	return out
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
			if p > 1 && p+1 == num {
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
