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

  ws, err := websocket.Dial(url, "", dialOrigin)
  defer ws.Close()
  if err != nil {
    fmt.Println("Error:", err)
  }
  fmt.Println("OK")

  for {
    var msg = make([]byte, 512)
    var n int
    if n, err = ws.Read(msg); err != nil {
      fmt.Println(err)
    } else if n > 0 {
      fmt.Printf("Received: %s.\n", msg[:n])
    }

  }
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
