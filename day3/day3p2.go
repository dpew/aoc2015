package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Position struct{ X, Y int }

func advance(dir byte, pos Position) Position {
	switch dir {
	case '<':
		pos = Position{X: pos.X - 1, Y: pos.Y}
	case '>':
		pos = Position{X: pos.X + 1, Y: pos.Y}
	case '^':
		pos = Position{X: pos.X, Y: pos.Y - 1}
	case 'v':
		pos = Position{X: pos.X, Y: pos.Y + 1}
	}
	return pos
}

func main() {
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	count := make(map[Position]int)
	var pos [2]Position
	index := 0

	count[pos[index]]++
	for _, c := range contents {
		pos[index] = advance(c, pos[index])
		count[pos[index]]++
		index = (index + 1) % 2
	}

	var number int
	for _, v := range count {
		if v > 1 {
			number += 1
		}
	}
	fmt.Printf("Found %d of %d with one or more gift\n", len(count), len(contents))
	fmt.Printf("Found %d of %d with more than one gift\n", number, len(contents))
}
