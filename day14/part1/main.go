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
	reaction         map[string]int
	reactionsOrdered []string
	amount           int
	name             string
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
		resourceStruct := resourceElem{amount: d, name: r, reaction: make(map[string]int), reactionsOrdered: make([]string, 0)}
		nanofactory[resourceStruct.name] = 0
		for _, react := range reactionSplit {
			var (
				r string
				d int
			)
			fmt.Sscanf(string(react), "%d %s", &d, &r)
			resourceStruct.reaction[r] = d
			resourceStruct.reactionsOrdered = append(resourceStruct.reactionsOrdered, r)
		}
		// order the reactions by name

		// ORE will neve appear on the right side because it cannot be generated.
		resources[resourceStruct.name] = resourceStruct
	}
	generateResource("FUEL", 1)
}

func generateResource(r string, n int) {
	//fmt.Println(nanofactory)
	if r == "ORE" {
		nanofactory["ORE"] += n
		totalORE += n
	}
	if nanofactory[r] >= n {
		// We are done, we generated enough of this resource
		return
	}

	resource := resources[r]

	// Something about resource generation, or things getting not calculated.

	// If there aren't enough resources we generate them recursively.
	for k, v := range resource.reaction {
		for nanofactory[k] < v { // Amig nincs r-bol n mennyisegu addig kell futtatni? -> ennek kint kene lennie?
			generateResource(k, v)
		}
		//if nanofactory[k] < v {
		//}
	}
	// if there are enough resources, we deduct that amount that the forumla needs
	// and generate the resource.amount needed in the system.
	for k, v := range resource.reaction {
		//if nanofactory[k]-v < 0 {
		//	generateResource(k, -(nanofactory[k] - v)) // nope ended up being more
		//}
		nanofactory[k] -= v // Nem mehetne 0 ala... If it goes beyond 0 generate as much as it doesn't got below it
		// Also the problem is that it was able to geenerate 1 FUEL with on 11 ORE
	}
	// Generate `amount` resources.
	nanofactory[resource.name] += resource.amount
}
