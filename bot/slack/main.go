package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

const dialOrigin string = "http://localhost/"

var apiToken = os.Getenv("SLACK_API_TOKEN")

type botDetails struct {
	URL string
}

type SlackBot struct {
	//Events chan Event
}

func NewSlackBot() *SlackBot {
	return new(SlackBot)
}

// Start will start slack bot routines.
func (sb *SlackBot) Start() {
	fmt.Println("slack bot goroutine")
	bot := rtmStart()
	ws, _ := establishWSConnection(bot.URL)

	defer ws.Close()

	for {
		data := new(Event)
		err := websocket.JSON.Receive(ws, data)
		if err != nil {
			fmt.Println("WS Codec.Receive", err)
		}
		fmt.Println(data)
	}
}

// establishWSConnection will establish WS connection to slack API.
func establishWSConnection(url string) (*websocket.Conn, error) {
	ws, err := websocket.Dial(url, "", dialOrigin)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	fmt.Println("WS connection established")
	return ws, nil
}

// rtmStart will request new rtm session from slack API.
func rtmStart() *botDetails {
	resp, httpErr := http.Get("https://slack.com/api/rtm.start?token=" + apiToken)
	defer resp.Body.Close()

	if httpErr != nil {
		fmt.Println("http get error")
	}

	body, readAllErr := ioutil.ReadAll(resp.Body)
	if readAllErr != nil {
		fmt.Println("ioutil error")
	}

	var data = new(botDetails)
	unmarshalErr := json.Unmarshal(body, data)

	if unmarshalErr != nil {
		fmt.Println("Unmarshal error")
	}
	return data
}
