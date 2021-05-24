package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	dictByID := make(map[int]game)
	for _, g := range games {
		dictByID[g.id] = g
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nlist : list all the games\nquit : quit game\nid N : queries a game by id\n")

		if !reader.Scan() {
			break
		}

		fmt.Println()

		option := strings.Fields(reader.Text())

		isQuit := doOperation(option, games, dictByID)

		if isQuit {
			return
		}
	}
}

func doOperation(option []string, games []game, dictByID map[int]game) bool {
	if len(option) == 0 {
		return false
	}

	switch option[0] {
	case "quit":
		fmt.Println("Application End!")
		return true

	case "list":
		for _, g := range games {
			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)
		}

	case "id":
		if len(option) != 2 {
			fmt.Println("wrong id")
			return false
		}

		id, err := strconv.Atoi(option[1])
		if err != nil {
			fmt.Println("wrong id")
			return false
		}

		g, ok := dictByID[id]
		if !ok {
			fmt.Println("sorry. i don't have the game")
		} else {
			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)
		}

	default:
		fmt.Println("Invalid input")
		return false
	}
	return false
}
