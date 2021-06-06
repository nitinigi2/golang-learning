package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func readUserInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	fmt.Print("Enter Username: ")
	userName, _ := reader.ReadString('\n')

	return email, userName
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connected to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	email, userName := readUserInput()

	go func() {
		var message Message
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			json.Unmarshal(msg, &message)
			log.Println(message.Username + "sent: " + message.Message)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

		if text == "" {
			continue
		}

		if text == "quit\r\n" {
			//<-interrupt
			err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
		}

		messagePtr := createMessage(text, email, userName)
		bytes, _ := json.Marshal(messagePtr)

		err := c.WriteMessage(websocket.TextMessage, []byte(bytes))
		if err != nil {
			log.Println("write:", err)
			return
		}
		fmt.Println("Message sent")
	}
}

func createMessage(text string, email string, userName string) *Message {
	text = strings.Trim(text, "/r/n")
	return &Message{
		Email:    email,
		Username: userName,
		Message:  text,
	}
}
