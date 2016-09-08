package main

import (
	"net/http"
	"github.com/imdoroshenko/go-steambot/models"
  "github.com/imdoroshenko/go-steambot/router"
  "encoding/json"
)

func main() {
  appRouter := new(router.Router)
  appRouter.AddActions(
    router.NewAction("GET", "/", wishList),
    router.NewAction("GET", "/foo", func (r http.ResponseWriter, rq *http.Request, p map[string]string) []byte {return []byte("foo")}),
    router.NewAction("POST", "/bar", func (r http.ResponseWriter, rq *http.Request, p map[string]string) []byte {return []byte("bar")}))
	http.ListenAndServe(":8080", appRouter)
}

func wishList(res http.ResponseWriter, req *http.Request, params map[string]string) []byte {
  p := &models.Player{SteamID:76561198019326316}
  p.UploadWishList()
  result, _ := json.Marshal(p.WishList)
	return result
}
