package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type node struct {
	v        []byte
	parent   *node
	children []*node
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	bytesArr := bytes.Split(content, []byte("\n"))
	nodes := make(map[string]*node)
	for _, b := range bytesArr {
		split := bytes.Split(b, []byte(")"))
		s1 := bytes.TrimSpace(split[0])
		s2 := bytes.TrimSpace(split[1])
		// check if it exists
		if v, ok := nodes[string(s1)]; ok {
			// let's hope that the input doesn't contain duplicates
			// a child will not have multiple parents
			if c, cok := nodes[string(s2)]; cok {
				c.parent = v
				v.children = append(v.children, c)
			} else {
				c = &node{v: s2}
				c.parent = v
				v.children = append(v.children, c)
				nodes[string(c.v)] = c
			}
		} else {
			p := &node{v: s1}
			if c, cok := nodes[string(s2)]; cok {
				// since I'm going sequentially, I have to retro fit existing child nodes.
				c.parent = p
				p.children = append(p.children, c)
			} else {
				c = &node{v: s2}
				c.parent = p
				p.children = append(p.children, c)
				nodes[string(c.v)] = c
			}
			nodes[string(s1)] = p
		}
	}
	findShortestPath(nodes)
}

func moves(n *node) []*node {
	more := make([]*node, 0)
	if n.parent != nil {
		more = append(more, n.parent)
	}
	more = append(more, n.children...)
	return more
}

// the root node should be COM
func findShortestPath(m map[string]*node) {
	var root *node
	for _, n := range m {
		if string(n.v) == "YOU" {
			root = n
			break
		}
	}
	if root == nil {
		return
	}
	goal := []byte("SAN")
	queue := make([]*node, 0)
	cameFrom := make(map[*node]*node)
	cameFrom[root] = nil
	queue = append(queue, root)
	var end *node
	var next *node
	for len(queue) > 0 {
		next, queue = queue[0], queue[1:]
		if bytes.Equal(next.v, goal) {
			end = next
			fmt.Println("found SAN")
			break
		}
		for _, child := range moves(next) {
			if _, ok := cameFrom[child]; !ok {
				cameFrom[child] = next
				queue = append(queue, child)
			}
		}
	}

	steps := 0
	current := end
	for current != root {
		steps++
		current = cameFrom[current]
	}
	// -2 because no start and end
	fmt.Println("steps: ", steps-2)
}
