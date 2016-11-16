package main

import (
	"encoding/json"
	"net/http"

	"github.com/imdoroshenko/go-steambot/bot/slack"
	"github.com/imdoroshenko/go-steambot/models"
	"github.com/imdoroshenko/go-steambot/router"
)

func init() {
	var slackBot = slack.NewSlackBot()
	go slackBot.Start()
}

func main() {
	appRouter := new(router.Router)
	appRouter.AddActions(
		router.NewAction("GET", "/wishlist/[playerID]", wishList),
		router.NewAction("GET", "/foo", foo),
		router.NewAction("POST", "/bar", bar))
	http.ListenAndServe(":8080", appRouter)
}

func bar(res http.ResponseWriter, req *http.Request, params map[string]string) []byte {
	return []byte("bar")
}

func foo(res http.ResponseWriter, req *http.Request, params map[string]string) []byte {
	return []byte("foo")
}

func wishList(res http.ResponseWriter, req *http.Request, params map[string]string) []byte {
	p := models.NewPlayer(params["playerID"])
	p.UploadWishList()
	result, _ := json.Marshal(p.WishList)
	return result
}
