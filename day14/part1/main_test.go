package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestBasic(t *testing.T) {
	content, _ := ioutil.ReadFile("test.txt")
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	lines := bytes.Split(content, []byte("\n"))
	manifacture(lines)
	if totalORE != 31 {
		t.Fatal("Total ore: ", totalORE)
	}
}

func Test2Basic(t *testing.T) {
	totalORE = 0
	content, _ := ioutil.ReadFile("test2.txt")
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	lines := bytes.Split(content, []byte("\n"))
	manifacture(lines)
	if totalORE != 165 {
		t.Fatal("Total ore: ", totalORE)
	}
}

func Test3Basic(t *testing.T) {
	totalORE = 0
	content, _ := ioutil.ReadFile("test3.txt")
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	lines := bytes.Split(content, []byte("\n"))
	manifacture(lines)
	if totalORE != 13312 {
		t.Fatal("Total ore: ", totalORE)
	}
}

func Test4Basic(t *testing.T) {
	totalORE = 0
	content, _ := ioutil.ReadFile("test4.txt")
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	lines := bytes.Split(content, []byte("\n"))
	manifacture(lines)
	if totalORE != 180697 {
		t.Fatal("Total ore: ", totalORE)
	}
}

func Test5Basic(t *testing.T) {
	totalORE = 0
	content, _ := ioutil.ReadFile("test5.txt")
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	lines := bytes.Split(content, []byte("\n"))
	manifacture(lines)
	if totalORE != 2210736 {
		t.Fatal("Total ore: ", totalORE)
	}
}
