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
	iterations := 1000
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
	fmt.Println(runPart1(moons, iterations))
}

func generatePairs(moons []*moon) [][]*moon {
	pairs := make([][]*moon, 0)
	for i := 0; i < len(moons); i++ {
		for j := 0; j < len(moons); j++ {
			if moons[i].id == moons[j].id {
				continue
			}
			pair := []*moon{moons[i], moons[j]}
			// we don't insert 1,2 -> 2,1 because our gravity will apply the effect to the other pair as well
			if !containsReversePair(pairs, pair) {
				pairs = append(pairs, pair)
			}
		}
	}
	return pairs
}

func containsReversePair(pairs [][]*moon, pair []*moon) bool {
	for _, p := range pairs {
		if pair[0].id == p[1].id && pair[1].id == p[0].id {
			return true
		}
	}
	return false
}

// runPart1 returns the total energy in the system.
func runPart1(moons []*moon, iterations int) int {
	for i := 0; i < iterations; i++ {
		for i := 0; i < len(moons); i++ {
			for j := i; j < len(moons); j++ {
				moons[i].applyGravity(moons[j])
			}
		}

		for _, m := range moons {
			m.applyVelocity()
			//fmt.Println(m)
		}
		//fmt.Printf("============= Iteration: %d =============\n", i)
	}
	sum := 0

	for _, m := range moons {
		sum += m.calculateTotalEnergy()
	}

	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
