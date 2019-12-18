package main

import (
	"fmt"
	"reflect"
)

type position struct {
	x, y, z int
}

type moon struct {
	id int
	p  position
	v  position
}

func (m *moon) applyGravity(pair *moon) {
	if m.p.x < pair.p.x {
		m.v.x++
		pair.v.x--
	} else if m.p.x > pair.p.x {
		m.v.x--
		pair.v.x++
	}
	if m.p.y < pair.p.y {
		m.v.y++
		pair.v.y--
	} else if m.p.y > pair.p.y {
		m.v.y--
		pair.v.y++
	}
	if m.p.z < pair.p.z {
		m.v.z++
		pair.v.z--
	} else if m.p.z > pair.p.z {
		m.v.z--
		pair.v.z++
	}
}

func (m *moon) applyVelocity() {
	m.p.x += m.v.x
	m.p.y += m.v.y
	m.p.z += m.v.z
}

func (m moon) String() string {
	return fmt.Sprintf("pos<x: %d, y: %d, z: %d>, vel<x: %d, y: %d, z: %d>", m.p.x, m.p.y, m.p.z, m.v.z, m.v.y, m.v.z)
}

func (m *moon) calculateTotalEnergy() int {
	potential := abs(m.p.x) + abs(m.p.y) + abs(m.p.z)
	kinetic := abs(m.v.x) + abs(m.v.y) + abs(m.v.z)
	return potential * kinetic
}

func main() {
	/*
		<x=-15, y=1, z=4>
		<x=1, y=-10, z=-8>
		<x=-5, y=4, z=9>
		<x=4, y=6, z=-2>
	*/
	moons := []*moon{
		{0, position{-15, 1, 4}, position{0, 0, 0}},
		{1, position{1, -10, -8}, position{0, 0, 0}},
		{2, position{-5, 4, 9}, position{0, 0, 0}},
		{3, position{4, 6, -2}, position{0, 0, 0}},
	}
	fmt.Println(runPart2(moons))
}

type sums struct {
	xSteps, ySteps, zSteps int
	startMoon              moon
}

func (s sums) String() string {
	return fmt.Sprintf("xSteps: %d; ySteps: %d, zSteps: %d", s.xSteps, s.ySteps, s.zSteps)
}

type ExtractFunc func(m *moon) []int

func snapshot(moons []*moon, extractor ExtractFunc) [][]int {
	results := make([][]int, 0)
	for _, m := range moons {
		results = append(results, extractor(m))
	}
	return results
}

type Initial struct {
	Value   [][]int
	Shot    ExtractFunc
	Done    bool
	Counter int
}

func NewInitial(moons []*moon, shot ExtractFunc) *Initial {
	return &Initial{
		Value:   snapshot(moons, shot),
		Shot:    shot,
		Done:    false,
		Counter: 0,
	}
}

func (c *Initial) Compare(moons []*moon) bool {
	// we can say it's false to avoid calling SetDone again
	// it's not cheating because statistically it's not true
	// .... most of the time XD
	if c.Done {
		return false
	}
	return reflect.DeepEqual(snapshot(moons, c.Shot), c.Value)
}

func (c *Initial) SetDone(counter int) {
	c.Counter = counter
	c.Done = true
}

func allDone(initials []*Initial) bool {
	for _, initial := range initials {
		if !initial.Done {
			return false
		}
	}
	return true
}

// runPart1 returns the total energy in the system.
func runPart2(moons []*moon) int {
	initials := []*Initial{
		NewInitial(moons, func(m *moon) []int { return []int{m.p.x, m.v.x} }),
		NewInitial(moons, func(m *moon) []int { return []int{m.p.y, m.v.y} }),
		NewInitial(moons, func(m *moon) []int { return []int{m.p.z, m.v.z} }),
	}

	steps := 0

	for {
		steps++

		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				moons[i].applyGravity(moons[j])
			}
		}

		for _, m := range moons {
			m.applyVelocity()
		}

		for _, initial := range initials {
			if initial.Compare(moons) {
				initial.SetDone(steps)
			}
		}

		if allDone(initials) {
			break
		}

	}

	counters := make([]int, 0)
	for _, initial := range initials {
		counters = append(counters, initial.Counter)
	}

	return lcm(counters...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(x ...int) int {
	if len(x) == 0 {
		return 0
	} else if len(x) == 2 {
		for x[1] != 0 {
			x[0], x[1] = x[1], x[0]%x[1]
		}
	} else if len(x) > 2 {
		return gcd(x[0], gcd(x[1:]...))
	}
	return abs(x[0])
}

func lcm(x ...int) int {
	if len(x) > 2 {
		return lcm(x[0], lcm(x[1:]...))
	} else if x[0] == 0 && x[1] == 0 {
		return 0
	}
	return abs(x[0]*x[1]) / gcd(x[0], x[1])
}
