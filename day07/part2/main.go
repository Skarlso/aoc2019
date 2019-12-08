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

const (
	add = iota + 1
	multi
	input
	output
	jmp
	jmpf
	less
	eq
)

type machine struct {
	position   int
	memory     map[int]int
	input      []int
	output     int
	inputIndex int
	inc        int
}

func (m machine) String() string {
	return fmt.Sprintf("position: %d; memory: %+v; input: %+v; output: %d; inputIndex: %d; inc: %d",
		m.position,
		m.memory,
		m.input,
		m.output,
		m.inputIndex,
		m.inc)
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

	sequence := []int{0, 1, 2, 3, 4}
	sequences := permutation(sequence)

	var max int
	for _, seq := range sequences {
		m := memClone(memory)
		a := &machine{memory: m, inc: 4}
		m = memClone(memory)
		b := &machine{memory: m, inc: 4}
		m = memClone(memory)
		c := &machine{memory: m, inc: 4}
		m = memClone(memory)
		d := &machine{memory: m, inc: 4}
		m = memClone(memory)
		e := &machine{memory: m, inc: 4}

		a.input = []int{seq[0]}
		b.input = []int{seq[1]}
		c.input = []int{seq[2]}
		d.input = []int{seq[3]}
		e.input = []int{seq[4]}
		var (
			allDone bool
			out     int
			eout    int
		)
		for !allDone {
			a.input = append(a.input, eout)
			out, allDone = a.processProgram()
			//fmt.Println("a: ", a)
			b.input = append(b.input, out)
			out, allDone = b.processProgram()
			//fmt.Println("b: ", b)
			c.input = append(c.input, out)
			out, allDone = c.processProgram()
			//fmt.Println("c: ", c)
			d.input = append(d.input, out)
			out, allDone = d.processProgram()
			//fmt.Println("d: ", d)
			e.input = append(e.input, out)
			eout, allDone = e.processProgram()
			//fmt.Println("e: ", e)
			if eout > max {
				max = eout
			}
		}
		//time.Sleep(1 * time.Second)
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

func (m *machine) processProgram() (out int, done bool) {
loop:
	for {
		opcode := m.memory[m.position]
		op, modes := getOpCodeAndModes(opcode)
		//fmt.Println(memory)
		//time.Sleep(1 * time.Second)
		//fmt.Println("i, op: ", i, op)
		switch op {
		case add:
			args := getArguments(3, m.position, modes, m.memory)
			m.memory[args[2]] = args[0] + args[1]
			m.inc = 4
		case multi:
			args := getArguments(3, m.position, modes, m.memory)
			m.memory[args[2]] = args[0] * args[1]
			m.inc = 4
		case input:
			m.memory[m.memory[m.position+1]] = m.input[m.inputIndex]
			m.inc = 2
			m.inputIndex++
		case output:
			var oout int
			if len(modes) > 0 {
				switch modes[0] {
				case position:
					oout = m.memory[m.memory[m.position+1]]
				case immediate:
					oout = m.memory[m.position+1]
				}
			} else {
				oout = m.memory[m.memory[m.position+1]]
			}
			out = oout
			m.inc = 2
			m.position += m.inc
			return out, false
		case jmp:
			args := getArguments(2, m.position, modes, m.memory)
			if args[0] != 0 {
				m.position = args[1]
				m.inc = 0
			} else {
				m.inc = 3
			}
			//fmt.Printf("5 i: %d args: %+v\n", i, args)
		case jmpf:
			args := getArguments(2, m.position, modes, m.memory)
			if args[0] == 0 {
				m.position = args[1]
				m.inc = 0
			} else {
				m.inc = 3
			}
			//fmt.Printf("6 i: %d args: %+v\n", i, args)
		case less:
			args := getArguments(3, m.position, modes, m.memory)
			if args[0] < args[1] {
				m.memory[args[2]] = 1
			} else {
				m.memory[args[2]] = 0
			}
			//fmt.Printf("7 i: %d args: %+v\n", i, args)
			m.inc = 4
		case eq:
			args := getArguments(3, m.position, modes, m.memory)
			if args[0] == args[1] {
				m.memory[args[2]] = 1
			} else {
				m.memory[args[2]] = 0
			}
			//fmt.Printf("8 i: %d args: %+v\n", i, args)
			m.inc = 4
		case 99:
			break loop
		}
		m.position += m.inc
	}

	return out, true
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
