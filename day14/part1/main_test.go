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
}
