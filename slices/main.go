package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// var todo []string

	// todo = append(todo, "eat")
	// todo = append(todo, "sleep")
	// todo = append(todo, "code")
	// todo = append(todo, "repeat")

	// s.Show("todo", todo)

	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	s.MaxPerLine = 4

	s.Show("All items", items)

	top3 := items[:3]
	s.Show("Top 3 items", top3)

	l := len(items)

	// you can use variables in a slice expression
	last4 := items[l-4:]
	s.Show("Last 4 items", last4)

	// reslicing: slicing another sliced slice
	mid := last4[1:3]
	s.Show("Last4[1:3]", mid)

	// the same elements can be in different indexes
	// fmt.Println(items[9], last4[0])

	// slicing returns a slice with the same type of the sliced slice
	fmt.Printf("slicing : %T %[1]q\n", items[2:3]) // slicing : []string ["tetris"]

	// indexing returns a single element with the type of the indexed slice's element type
	fmt.Printf("indexing: %T %[1]q\n", items[2]) // indexing: string "tetris"
}
