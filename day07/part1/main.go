package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Skarlso/intcode"
)

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
		machineA.Position = 0
		out, _ := machineA.ProcessProgram()
		m = memClone(memory)
		machineB.Memory = m
		machineB.Input = []int{seq[1], out[0]}
		machineB.Position = 0
		out, _ = machineB.ProcessProgram()
		m = memClone(memory)
		machineC.Memory = m
		machineC.Input = []int{seq[2], out[0]}
		machineC.Position = 0
		out, _ = machineC.ProcessProgram()
		m = memClone(memory)
		machineD.Memory = m
		machineD.Input = []int{seq[3], out[0]}
		machineD.Position = 0
		out, _ = machineD.ProcessProgram()
		m = memClone(memory)
		machineE.Memory = m
		machineE.Input = []int{seq[4], out[0]}
		machineE.Position = 0
		out, _ = machineE.ProcessProgram()
		if out[0] > max {
			max = out[0]
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
