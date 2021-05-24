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

func unMarshallData() ([]game, error) {
	var unmarshalledData []jsonGame
	if err := json.Unmarshal([]byte(data), &unmarshalledData); err != nil {
		return []game{}, err
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
	return games, nil
}

func main() {

	games, err := unMarshallData()
	if err != nil {
		fmt.Println(err)
		return
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
