package main

import (
	"fmt"
)

type item struct {
	item_name   string
	item_type   string
	hp_modifier int
}

type character struct {
	playername      string
	playerhp        int
	equipped_weapon item
	equipped_armor  item
	inventory       []item
}

func create_player(name string) character {
	var c character
	c.playername = name
	c.playerhp = 100
	hp_potion := create_item("HP", "Potion", 20)
	c.inventory = append(c.inventory, hp_potion, hp_potion, hp_potion)
	c.equipped_armor = create_item("Beginner", "Armor", 20)
	c.equipped_weapon = create_item("Beginner", "Sword", 5)

	return c
}

func create_item(it_name string, it_type string, hp_mod int) item {
	var i item
	i.item_name = it_name
	i.item_type = it_type
	i.hp_modifier = hp_mod

	return i
}

func hp_change_char(c *character, hp_change int) {
	c.playerhp += hp_change

	if c.playerhp >= 100 {
		c.playerhp = 100
	}
}

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func main() {
	fmt.Println("Insert player name: ")
	var nameinput string
	fmt.Scan(&nameinput)
	var player = create_player(nameinput)

	fmt.Println("Player Name: ", player.playername)
	fmt.Println("Player HP: ", player.playerhp)
	fmt.Println("Player Equipped Weapon: ", player.equipped_weapon)
	fmt.Println("Player Equipped Armor: ", player.equipped_armor)
	fmt.Println("Player Inventory: ", player.inventory)

	for {
		fmt.Println("Pick where you want to go")
		fmt.Println("1. The Forest (lvl 1 area)")
		fmt.Println("2. The Mountain (lvl 2 area)")
		fmt.Println("3. The Dungeon (lvl 3 area)")
		fmt.Println("4. The Shop")
		fmt.Println("5. The Tavern (heal to full health)")

		var pin int
		fmt.Scan(&pin)

		switch pin {
		case 1:
			fmt.Println("Coming Soon")
		case 2:
			fmt.Println("Coming Soon")
		case 3:
			fmt.Println("Coming Soon")
		case 4:
			fmt.Println("Coming Soon")
		case 5:
			hp_change_char(&player, 100)
			fmt.Println("Fully Healed")
			fmt.Println(player.playerhp)
		}

	}
}

// changed := strconv.Itoa(player.playerhp)

// toFile := []byte(player.playername + " " + changed)
// err := os.WriteFile("text.txt", toFile, 0644)
// check(err)
