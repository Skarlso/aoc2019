package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), ",")
	mem := make([]int, len(split))
	pc := 0
	base := 0

	for i, s := range split {
		mem[i], _ = strconv.Atoi(s)
	}

	for {
		ins := fmt.Sprintf("%05d", mem[pc])
		op, _ := strconv.Atoi(ins[3:])
		arg := func(i int) (addr int) {
			switch ins[3-i] {
			case '0':
				addr = mem[pc+i]
			case '1':
				addr = pc + i
			case '2':
				addr = base + mem[pc+i]
			}
			if len(mem) <= addr {
				mem = append(mem, make([]int, addr-len(mem)+1)...)
			}
			return
		}

		switch op {
		case 1:
			mem[arg(3)] = mem[arg(1)] + mem[arg(2)]
		case 2:
			mem[arg(3)] = mem[arg(1)] * mem[arg(2)]
		case 3:
			fmt.Scan(&mem[arg(1)])
		case 4:
			fmt.Println(mem[arg(1)])
		case 5:
			if mem[arg(1)] != 0 {
				pc = mem[arg(2)]
				continue
			}
		case 6:
			if mem[arg(1)] == 0 {
				pc = mem[arg(2)]
				continue
			}
		case 7:
			if mem[arg(1)] < mem[arg(2)] {
				mem[arg(3)] = 1
			} else {
				mem[arg(3)] = 0
			}
		case 8:
			if mem[arg(1)] == mem[arg(2)] {
				mem[arg(3)] = 1
			} else {
				mem[arg(3)] = 0
			}
		case 9:
			base += mem[arg(1)]
		case 99:
			return
		}

		pc += []int{1, 4, 4, 2, 2, 3, 3, 4, 4, 2}[op]
	}
}
