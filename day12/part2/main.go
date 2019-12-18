package main

import "fmt"

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

// runPart1 returns the total energy in the system.
func runPart2(moons []*moon) int {
	// steps := 0
	startM1 := *moons[0]
	startM2 := *moons[1]
	startM3 := *moons[2]
	startM4 := *moons[3]

	endXStep := 0
	endYStep := 0
	endZStep := 0

	steps := 0

	for {
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				moons[i].applyGravity(moons[j])
			}
		}

		for _, m := range moons {
			m.applyVelocity()
		}

		if moons[0].p.x == startM1.p.x &&
			moons[0].v == startM1.v &&
			moons[1].p.x == startM2.p.x &&
			moons[1].v == startM2.v &&
			moons[2].p.x == startM3.p.x &&
			moons[2].v == startM3.v &&
			moons[3].p.x == startM4.p.x &&
			moons[3].v == startM4.v &&
			endXStep == 0 {
			endXStep = steps
		}
		if moons[0].p.y == startM1.p.y &&
			moons[0].v == startM1.v &&
			moons[1].p.y == startM2.p.y &&
			moons[1].v == startM2.v &&
			moons[2].p.y == startM3.p.y &&
			moons[2].v == startM3.v &&
			moons[3].p.y == startM4.p.y &&
			moons[3].v == startM4.v &&
			endYStep == 0 {
			endYStep = steps
		}

		if moons[0].p.z == startM1.p.z &&
			moons[0].v == startM1.v &&
			moons[1].p.z == startM2.p.z &&
			moons[1].v == startM2.v &&
			moons[2].p.z == startM3.p.z &&
			moons[2].v == startM3.v &&
			moons[3].p.z == startM4.p.z &&
			moons[3].v == startM4.v &&
			endZStep == 0 {
			endZStep = steps
		}
		steps++

		if endXStep > 0 && endYStep > 0 && endZStep > 0 {
			break
		}
	}

	fmt.Println(lcm(endXStep, endYStep, endZStep))
	return 0
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
