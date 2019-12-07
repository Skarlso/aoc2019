package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type node struct {
	v []byte
	parent *node
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
		// check if it exists
		if v, ok := nodes[string(split[0])]; ok {
			// let's hope that the input doesn't contain duplicates
			// a child will not have multiple parents
			if c, cok := nodes[string(split[1])]; cok {
				// since I'm going sequentially, I have to retro fit existing child nodes.
				c.parent = v
			} else {
				c = &node{v: split[1]}
				c.parent = v
				v.children = append(v.children, c)
				nodes[string(c.v)] = c
			}
		} else {
			p := &node{v: split[0]}
			if c, cok := nodes[string(split[1])]; cok {
				// since I'm going sequentially, I have to retro fit existing child nodes.
				c.parent = p
			} else {
				c = &node{v: split[1]}
				c.parent = p
				p.children = append(p.children, c)
				nodes[string(c.v)] = c
			}
			nodes[string(split[0])] = p
		}
	}
	travers(nodes)
}

func countBack(n *node) int {
	var sum int
	for n != nil {
		if n.parent != nil {
			sum++
		}
		n = n.parent
	}
	return sum
}

// the root node should be COM
func travers(m map[string]*node) {
	var root *node
	for _, n := range m {
		if string(n.v) == "COM" {
			root = n
			break
		}
	}
	var sum int
	queue := make([]*node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		sum += countBack(nextUp)
		//fmt.Println("node: ", string(nextUp.v))
		//if nextUp.parent != nil {
		//	fmt.Println("Parent: ", string(nextUp.parent.v))
		//}
		//fmt.Println("------")
		if len(nextUp.children) > 0 {
			//fmt.Printf("Number of children for node %q is %d\n", string(nextUp.v), len(nextUp.children))
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	fmt.Println("number of connections: ", sum)
}


