package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Skarlso/intcode"
)

type point struct {
	x, y int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename")
	}
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
	m := intcode.NewMachine(memory)
	// m.Input = []int{0}
	var (
		out  []int
		done bool
		// grid = make(map[point]int)
		// x, y int
	)
	reader := bufio.NewReader(os.Stdin)

	provideInput := true
	for !done {
		// While not done, keep displaying what's going on and display. Add the newline explicitly.
		out, done = m.ProcessProgram()
		display(out)
		if provideInput {
			text, _ := reader.ReadString('\n')
			ascii := make([]int, 0)
			for _, c := range text {
				ascii = append(ascii, int(c))
			}
			m.Input = ascii
		}
		if strings.Contains(convertToString(out), "Continuous video feed?") {
			provideInput = false
		}
	}
}

func convertToString(out []int) (s string) {
	for _, v := range out {
		s += string(v)
	}
	return
}

func display(out []int) {
	for _, v := range out {
		fmt.Print(string(v))
	}
}
