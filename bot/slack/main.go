package slack

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "golang.org/x/net/websocket"
  "os"
)

type JSONResponse map[string]interface{}

const dialOrigin string = "http://localhost/"

var apiToken string = os.Getenv("SLACK_API_TOKEN")

var botDetails JSONResponse

func Start () {
  fmt.Println("slack bot goroutine")

  botDetails = RTMStart()
  url, ok := botDetails["url"].(string)
  if (ok == false) {
    fmt.Println("url not string")
  }

  ws, _ := EstablishWSConnection(url)
  defer ws.Close()

  for {
    var data string
    websocket.Message.Receive(ws, &data)
    fmt.Printf("Received: %s\n", data)
  }
}

func EstablishWSConnection (url string) (*websocket.Conn, error) {
  ws, err := websocket.Dial(url, "", dialOrigin)
  if err != nil {
    fmt.Println("Error:", err)
    return nil, err
  }
  fmt.Println("WS connection established")
  return ws, nil
}

func RTMStart () JSONResponse {
  resp, httpErr := http.Get("https://slack.com/api/rtm.start?token=" + apiToken)
  defer resp.Body.Close()

  if httpErr != nil {
    fmt.Println("http get error")
    //return nil, err;
  }

  body, readAllErr := ioutil.ReadAll(resp.Body)
  if readAllErr != nil {
    fmt.Println("ioutil error")
  }

  var botDetails JSONResponse
  unmarshalErr := json.Unmarshal(body, &botDetails)

  if unmarshalErr != nil {
    fmt.Println("Unmarshal error")
  }
  return botDetails
}
