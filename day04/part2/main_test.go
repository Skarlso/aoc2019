package main

import "testing"

func TestBasic(t *testing.T) {
	b := check(111111)
	if b != false {
		t.Fatal("fuck")
	}
	b = check(223450)
	if b != false {
		t.Fatal("fuck")
	}
	b = check(123789)
	if b != false {
		t.Fatal("fuck")
	}
	b = check(111122)
	if b != true {
		t.Fatal("fuck")
	}
	b = check(112233)
	if b != true {
		t.Fatal("fuck")
	}
	b = check(123444)
	if b != false {
		t.Fatal("fuck")
	}
}
