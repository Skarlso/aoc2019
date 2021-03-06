package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	iterations := 10
	moons := []*moon{
		{0, position{-1, 0, 2}, position{0, 0, 0}},
		{1, position{2, -10, -7}, position{0, 0, 0}},
		{2, position{4, -8, 8}, position{0, 0, 0}},
		{3, position{3, 5, -1}, position{0, 0, 0}},
	}
	sum := runPart1(moons, iterations)
	if sum != 179 {
		t.Fatal("Sum was: ", sum)
	}
}

func Test2Basic(t *testing.T) {
	iterations := 100
	/*
	<x=-8, y=-10, z=0>
	<x=5, y=5, z=10>
	<x=2, y=-7, z=3>
	<x=9, y=-8, z=-3>
	 */
	moons := []*moon{
		{0, position{-8, -10, 0}, position{0, 0, 0}},
		{1, position{5, 5, 10}, position{0, 0, 0}},
		{2, position{2, -7, 3}, position{0, 0, 0}},
		{3, position{9, -8, -3}, position{0, 0, 0}},
	}
	sum := runPart1(moons, iterations)
	if sum != 1940 {
		t.Fatal("Sum was: ", sum)
	}
}
