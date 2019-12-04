package main

import "fmt"

func main() {
	from := 264360
	to := 746325
	n := 0
	for i := from; i <= to; i++ {
		if check(i) {
			n++
		}
	}

	fmt.Println(n)
}

func check(a int) bool {
	digits := make([]int, 0)
	for a > 0 {
		n := a % 10
		digits = append(digits, n)
		a /= 10
	}
	double := false
	for i := 0; i < len(digits); i++ {
		if i+1 < len(digits) {
			if digits[i] == digits[i+1] {
				double = true
			}
			if digits[i] < digits[i+1] {
				return false
			}
		}
	}
	return double
}
