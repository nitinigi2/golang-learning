package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {

	fmt.Println("Hello World Emoji!")

	emoji.Println(":taco: Taco!!!")

	pizzaMessage := emoji.Sprint("I like a :kiwi: and :sushi:!!")
	fmt.Println(pizzaMessage)
}
