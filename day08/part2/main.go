package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file := os.Args[1]
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	wide := 25
	tall := 6
	image := make([][]rune, tall)
	for i := 0; i < len(image); i++ {
		image[i] = make([]rune, wide)
		for j := 0; j < wide; j++ {
			image[i][j] = '.'
		}
	}

	var rowCount int
	for i := 0; i < len(content); i += wide * tall {
		if i+(wide*tall) > len(content) {
			break
		}
		layer := string(content[i : i+(wide*tall)])
		for j := 0; j < len(layer); j += wide {
			if j+wide > len(layer) {
				break
			}
			r := layer[j : j+wide]
			for k, v := range r {
				if image[rowCount][k] == '0' || image[rowCount][k] == '1' {
					continue
				}
				if image[rowCount][k] == '2' || image[rowCount][k] == '.' {
					image[rowCount][k] = v
				}
			}
			rowCount = (rowCount + 1) % tall
		}
	}
	fmt.Println("And the image is.... :")
	display(image)
}

func display(image [][]rune) {
	// Display the image.
	for _, col := range image {
		for _, row := range col {
			if row == '1' {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
