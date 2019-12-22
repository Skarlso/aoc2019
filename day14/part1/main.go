package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// determines how many resources there are in the system.
var nanofactory = make(map[string]int)
var resources = make(map[string]resourceElem)
var totalORE int

type resourceElem struct {
	reaction map[string]int
	amount   int
	name     string
}

func main() {
	filename := os.Args[1]
	content, _ := ioutil.ReadFile(filename)
	if content[len(content)-1] == 0x0a {
		content = content[:len(content)-1]
	}
	lines := bytes.Split(content, []byte("\n"))

	manifacture(lines)
	fmt.Println(totalORE)
}

func manifacture(lines [][]byte) {
	nanofactory["ORE"] = 0
	// The ore count in the system. If someone wants ORE we just increase this number. At the
	// end the result of this number is how much ORE we need in the system all together.
	for _, l := range lines {
		//fmt.Println(string(l))
		split := bytes.Split(l, []byte(" => "))
		reaction := split[0]
		resource := split[1]
		//_, _ = fmt.Sscanf(string(l), "%s => %s", &reaction, &resource)
		//fmt.Println(string(reaction))
		reactionSplit := bytes.Split(reaction, []byte(", "))

		var (
			r string
			d int
		)
		fmt.Sscanf(string(resource), "%d %s", &d, &r)
		resourceStruct := resourceElem{amount: d, name: r, reaction: make(map[string]int)}
		nanofactory[resourceStruct.name] = 0
		for _, react := range reactionSplit {
			var (
				r string
				d int
			)
			fmt.Sscanf(string(react), "%d %s", &d, &r)
			resourceStruct.reaction[r] = d
		}
		// ORE will neve appear on the right side because it cannot be generated.
		resources[resourceStruct.name] = resourceStruct
	}
	generateResource("FUEL", 1)
}

func generateResource(r string, n int) {
	if r == "ORE" {
		nanofactory["ORE"] += n
		totalORE += n
	}
	if nanofactory[r] >= n {
		// We are done, we generated enough of this resource
		return
	}

	resource := resources[r]
	// If there aren't enough resources we generate them recursively.
	for k, v := range resource.reaction {
		if nanofactory[k] < v {
			generateResource(k, v)
		}
	}
	nanofactory[resource.name] += resource.amount
	// if there are enough resources, we deduct that amount that the forumla needs
	// and generate the resource.amount needed in the system.
	for k, v := range resource.reaction {
		if nanofactory[k]-v > 0 {
			nanofactory[k] -= v
		} else {
			nanofactory[k] = 0
		}
	}
	// Generate `amount` resources.
}
