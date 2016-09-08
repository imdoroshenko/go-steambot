package models

import (
  "net/http"
  "github.com/PuerkitoBio/goquery"
  "fmt"
)

type Player struct {
  SteamID  string
  WishList []string
}

func (p *Player) UploadWishList()   {
  fmt.Println("UploadWishList", p.SteamID)
  resp, err := http.Get("http://steamcommunity.com/profiles/" + p.SteamID + "/wishlist/")
  defer resp.Body.Close()

  if err != nil {
    // handle error
    fmt.Println("Request error")
  }

  doc, err := goquery.NewDocumentFromResponse(resp)

  if err != nil {
    fmt.Println("goquery error")
  }

  doc.Find(".wishlistRow").Each(func(i int, s *goquery.Selection) {
    gameId, _ := s.Attr("id")
    p.WishList = append(p.WishList, gameId)
    fmt.Printf("Wishlist item %d: %s\n", i, gameId)
  })
}
