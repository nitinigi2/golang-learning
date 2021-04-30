package main

/*
	description : Use go modules, pick three random smiles from https://github.com/kyokomi/emoji and print them into the console
*/

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kyokomi/emoji/v2"
)

func main() {

	fmt.Println("Hello Emoji!")

	emoji.Println(":taco: Taco!!!")

	pizzaMessage := emoji.Sprint("I like a :kiwi: and :sushi:!!")
	fmt.Println(pizzaMessage)

	fmt.Println("\nRandom Emojis :")
	randomEmoji(3)
}

/*
	 noOfRandomEmoji - number of emoji's
     getting all available emojis and storing in slice
	 generating random indexes to pick emoji's from slice
	 print emojis
*/
func randomEmoji(noOfRandomEmoji int) {
	rand.Seed(time.Now().UTC().UnixNano()) // providing seed to generate random numbers always
	emogies := emoji.CodeMap()
	slice := make([]string, 0)
	for key := range emogies {
		slice = append(slice, key)
	}
	for noOfRandomEmoji > 0 {
		randomNumber := rand.Intn(len(slice))
		randomEmoji := emoji.Sprint(slice[randomNumber])
		fmt.Println(randomEmoji)
		noOfRandomEmoji--
	}
}
