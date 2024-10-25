package main

import (
	"fmt"
)

type player struct {
	playername string
	playerhp   int
	inventory  []string
}

func create_player(name string) player {
	var p player
	p.playername = name
	p.playerhp = 100
	p.inventory = append(p.inventory, "sword:weapon")

	return p
}

func main() {
	fmt.Println("Insert player name: ")
	var nameinput string
	fmt.Scan(&nameinput)
	var player = create_player(nameinput)

	fmt.Println(player)
}
