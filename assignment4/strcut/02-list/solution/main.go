package main

import (
	"bufio"
	"fmt"
	"os"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func loadGameData() []game {
	return []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{
			item:  item{id: 2, name: "x-com 2", price: 40},
			genre: "strategy",
		},
		{
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}
}

func main() {

	games := loadGameData()

	reader := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println("\nlist : list all the games\nquit : quit game\n")

		fmt.Println()

		if !reader.Scan() {
			break
		}

		option := reader.Text()

		isQuit := doOperation(option, games)

		if isQuit {
			return
		}

	}
}

func doOperation(option string, games []game) bool {
	switch option {
	case "quit":
		fmt.Println("Program Exit!")
		return true

	case "list":
		for _, g := range games {
			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)
		}
		return false

	default:
		fmt.Println("Invalid input")
		return false
	}
}
