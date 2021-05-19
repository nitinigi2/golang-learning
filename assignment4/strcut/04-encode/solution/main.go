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
	Id    int
	Name  string
	Price int
}

type game struct {
	item
	genre string
}

type jsonGame struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
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

	DictById := make(map[int]game)

	for _, g := range games {
		DictById[g.Id] = g
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
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.Id, g.Name, "("+g.genre+")", g.Price)
			}

		case "id":
			if len(option) != 2 {
				fmt.Println("wrong Id")
				continue
			}

			Id, err := strconv.Atoi(option[1])
			if err != nil {
				fmt.Println("wrong Id")
				continue
			}

			g, ok := DictById[Id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			} else {
				fmt.Printf("#%d: %-15q %-20s $%d\n", g.Id, g.Name, "("+g.genre+")", g.Price)
			}

		case "save":

			jsonData := make([]jsonGame, 0)

			for _, g := range games {
				jsonData = append(jsonData,
					jsonGame{g.Id, g.Name, g.genre, g.Price})
			}

			out, err := json.Marshal(jsonData)
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}

			fmt.Println(string(out))
			return

		default:
			fmt.Println("Invalid input")
		}
	}
}
