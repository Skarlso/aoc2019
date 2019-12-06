package main

import (
	"fmt"
	"testing"
)

func TestOpMode(t *testing.T) {
	op, modes := getOpCodeAndModes(1100)
	if op != 0 {
		t.Fatal("not 0: ", op)
	}
	fmt.Println(modes)

	op, modes = getOpCodeAndModes(1002)
	if op != 2 {
		t.Fatal("not 0: ", op)
	}
	fmt.Println(modes)
}
