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

	var max int
	for _, seq := range sequences {
		m := memClone(memory)
		a := intcode.NewMachine(m)
		m = memClone(memory)
		b := intcode.NewMachine(m)
		m = memClone(memory)
		c := intcode.NewMachine(m)
		m = memClone(memory)
		d := intcode.NewMachine(m)
		m = memClone(memory)
		e := intcode.NewMachine(m)

		a.Input = []int{seq[0]}
		b.Input = []int{seq[1]}
		c.Input = []int{seq[2]}
		d.Input = []int{seq[3]}
		e.Input = []int{seq[4]}
		var (
			allDone bool
			out     []int
		)
		eout := []int{0}
		for !allDone {
			a.Input = append(a.Input, eout...)
			out, allDone = a.ProcessProgram()
			//fmt.Println("a: ", a)
			b.Input = append(b.Input, out...)
			out, allDone = b.ProcessProgram()
			//fmt.Println("b: ", b)
			c.Input = append(c.Input, out...)
			out, allDone = c.ProcessProgram()
			//fmt.Println("c: ", c)
			d.Input = append(d.Input, out...)
			out, allDone = d.ProcessProgram()
			//fmt.Println("d: ", d)
			e.Input = append(e.Input, out...)
			eout, allDone = e.ProcessProgram()
			//fmt.Println("e: ", e)
			if len(eout) == 1 {
				if eout[0] > max {
					max = eout[0]
				}
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
