package main

import (
	"fmt"
	"strconv"
)

const input = "59773590431003134109950482159532121838468306525505797662142691007448458436452137403459145576019785048254045936039878799638020917071079423147956648674703093863380284510436919245876322537671069460175238260758289779677758607156502756182541996384745654215348868695112673842866530637231316836104267038919188053623233285108493296024499405360652846822183647135517387211423427763892624558122564570237850906637522848547869679849388371816829143878671984148501319022974535527907573180852415741458991594556636064737179148159474282696777168978591036582175134257547127308402793359981996717609700381320355038224906967574434985293948149977643171410237960413164669930"

func main() {
	basePattern := []int{0, 1, 0, -1}
	phaseCount := 100
	//
	expandPhase := func(n int, pattern []int) []int {
		newPatter := make([]int, 0)
		for _, p := range pattern {
			for i := 0; i < n; i++ {
				newPatter = append(newPatter, p)
			}
		}
		return newPatter
	}

	shiftPattern := func(pattern []int) []int {
		head := pattern[0]
		tail := pattern[1:]
		tail = append(tail, head)
		return tail
	}

	in := input
	sum := 0
	lastOut := ""
	for p := 0; p < phaseCount; p++ {
		newPattern := shiftPattern(basePattern)
		out := ""
		for i := 0; i < len(in); i++ {
			for j := 0; j < len(in); j++ {
				c := in[j]
				cint, _ := strconv.Atoi(string(c))
				times := cint * newPattern[(j%len(newPattern))]
				sum += times
			}
			outS := strconv.Itoa(abs(sum))
			out += string(outS[len(outS)-1])
			sum = 0
			newPattern = expandPhase(i+2, basePattern)
			newPattern = shiftPattern(newPattern)
		}
		lastOut = out[:8]
		in = out
	}
	fmt.Println(lastOut)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
