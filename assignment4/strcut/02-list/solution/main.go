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
			item:  item{Id: 1, Name: "god of war", Price: 50},
			genre: "action adventure",
		},
		{
			item:  item{Id: 2, Name: "x-com 2", Price: 40},
			genre: "strategy",
		},
		{
			item:  item{Id: 3, Name: "minecraft", Price: 20},
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

		switch reader.Text() {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}

		default:
			fmt.Println("Invalid input")
		}
	}
}
