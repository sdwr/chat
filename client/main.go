package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"encoding/json"

	"golang.org/x/net/websocket"
)

var (
	port = flag.String("port", "9000", "port used for ws connection")
)

type Message struct {
	Type string `json:"Type"`
	Text string `json:"Text"`
}

type Room struct {
	Teaser string `json:"Teaser"`
	Description string `json:Description"`
	Contents string `json:Contents"`
}

func main() {
	flag.Parse()

	// connect
	ws, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// receive
	var data []byte
	var m Message
	go func() {
		for {
			err := websocket.JSON.Receive(ws, &data)
			if err != nil {
				fmt.Println("Error receiving message: ", err.Error())
				break
			}
			err = json.Unmarshal(data, &m)
			fmt.Println("Message: ", m)
		}
	}()

	// send
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		m := Message{
			Type: "COMMAND",
			Text: text,
		}
		messageJson, err := json.Marshal(m)
		err = websocket.JSON.Send(ws, messageJson)
		if err != nil {
			fmt.Println("Error sending message: ", err.Error())
			break
		}
	}
}

// connect connects to the local chat server at port <port>
func connect() (*websocket.Conn, error) {
	return websocket.Dial(fmt.Sprintf("ws://localhost:%s", *port), "", mockedIP())
}

// mockedIP is a demo-only utility that generates a random IP address for this client
func mockedIP() string {
	var arr [4]int
	for i := 0; i < 4; i++ {
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(256)
	}
	return fmt.Sprintf("http://%d.%d.%d.%d", arr[0], arr[1], arr[2], arr[3])
}
