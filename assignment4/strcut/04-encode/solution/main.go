package main

import (
	"bufio"
	"encoding/json"
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

type jsonGame struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	Genre string `json:"genre"`
	price int    `json:"price"`
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

	DictById := make(map[int]game)

	for _, g := range games {
		DictById[g.id] = g
	}

	reader := bufio.NewScanner(os.Stdin)
	for {

		fmt.Println("\nlist : list all the games\nquit : quit game\nid N : queries a game by id\nsave : exports the data to json and quits")

		if !reader.Scan() {
			break
		}

		fmt.Println()

		option := strings.Fields(reader.Text())
		if len(option) == 0 {
			continue
		}

		switch option[0] {
		case "quit":
			fmt.Println("Application End")
			return

		case "list":
			printGameData(games)

		case "id":
			findGameById(option, DictById)

		case "save":
			saveGame(games)

			return

		default:
			fmt.Println("Invalid input")
		}
	}
}

func saveGame(games []game) {
	jsonData := make([]jsonGame, 0)

	for _, g := range games {
		jsonData = append(jsonData,
			jsonGame{g.id, g.name, g.genre, g.price})
	}

	out, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Sorry:", err)
		return
	}

	fmt.Println(string(out))
}

func findGameById(option []string, DictById map[int]game) {
	if len(option) != 2 {
		fmt.Println("wrong id")
		return
	}

	id, err := strconv.Atoi(option[1])
	if err != nil {
		fmt.Println("wrong id")
		return
	}

	g, ok := DictById[id]
	if !ok {
		fmt.Println("game id not found")
		return
	} else {
		fmt.Printf("#%d: %-15q %-20s $%d\n", g.id, g.name, "("+g.genre+")", g.price)
	}
}

func printGameData(games []game) {
	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.id, g.name, "("+g.genre+")", g.price)
	}
}
