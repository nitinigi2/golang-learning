package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

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
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

func main() {
	var unmarshalledData []jsonGame
	if err := json.Unmarshal([]byte(data), &unmarshalledData); err != nil {
		fmt.Println(err)
		return
	}

	var games []game
	for _, dg := range unmarshalledData {
		games = append(games,
			game{
				item:  item{dg.ID, dg.Name, dg.Price},
				genre: dg.Genre,
			},
		)
	}

	DictById := make(map[int]game)
	for _, g := range games {
		DictById[g.id] = g
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("list : list all the games\nquit : quit game\nid N : queries a game by id+ \nsave : exports the data to json and quits\n")

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
					g.id, g.name, "("+g.genre+")", g.price)
			}

		case "id":
			if len(option) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(option[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := DictById[id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)

		case "save":
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.Marshal(encodable)
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
