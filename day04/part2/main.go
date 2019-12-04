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
	for i := len(digits)/2 - 1; i >= 0; i-- {
		opp := len(digits) - 1 - i
		digits[i], digits[opp] = digits[opp], digits[i]
	}
	double := false
	for i := 0; i < len(digits); i++ {
		if i+1 >= len(digits) {
			break
		}
		if digits[i] > digits[i+1] {
			return false
		}

		if digits[i] == digits[i+1] {
			if i+2 < len(digits) && digits[i+2] != digits[i+1] {
				double = true
			} else if i+2 >= len(digits) {
				double = true
			}

			addToI := 0
			for j := i + 2; j < len(digits); j++ {
				if digits[j] == digits[i] {
					addToI++
					continue
				}
				break
			}
			i += addToI
		}
	}
	return double
}
