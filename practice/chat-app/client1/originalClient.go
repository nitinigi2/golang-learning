// package main

// import (
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/url"
// 	"os"
// 	"os/signal"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// var addr = flag.String("addr", "localhost:8080", "http service address")

// type Message struct {
// 	Email    string `json:"email"`
// 	Username string `json:"username"`
// 	Message  string `json:"message"`
// }

// func main() {
// 	flag.Parse()
// 	log.SetFlags(0)

// 	interrupt := make(chan os.Signal, 1)
// 	signal.Notify(interrupt, os.Interrupt)

// 	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
// 	log.Printf("connecting to %s", u.String())

// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer c.Close()

// 	done := make(chan struct{})

// 	go func() {
// 		defer close(done)
// 		for {
// 			_, message, err := c.ReadMessage()
// 			if err != nil {
// 				log.Println("read:", err)
// 				return
// 			}
// 			log.Printf("recv: %s", message)
// 		}
// 	}()

// 	sampleMessage := Message{
// 		Email:    "abc.gmail.com",
// 		Username: "abc",
// 		Message:  "message 1",
// 	}
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	bytes, _ := json.Marshal(sampleMessage)

// 	fmt.Println(bytes)

// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case <-ticker.C:
// 			err := c.WriteMessage(websocket.TextMessage, []byte(bytes))
// 			if err != nil {
// 				log.Println("write:", err)
// 				return
// 			}
// 		case <-interrupt:
// 			log.Println("interrupt")

// 			// Cleanly close the connection by sending a close message and then
// 			// waiting (with timeout) for the server to close the connection.
// 			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
// 			if err != nil {
// 				log.Println("write close:", err)
// 				return
// 			}
// 			select {
// 			case <-done:
// 			case <-time.After(time.Second):
// 			}
// 			return
// 		}
// 	}
// }
