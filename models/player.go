package models

import (
  "fmt"
  "github.com/imdoroshenko/go-steambot/loader"
  "github.com/PuerkitoBio/goquery"
  "strings"
)

func NewPlayer(SteamID string) *Player {
  p := new(Player)
  p.SteamID = SteamID
  return p
}

type Player struct {
  SteamID  string
  WishList []*App
}

// wishlist is not availible in steam web API,
// so I upload players wishlist from steam profile web page
func (p *Player) UploadWishList()   {
  fmt.Println("UploadWishList", p.SteamID)
  doc, _ := loader.Get("http://steamcommunity.com/profiles/" + p.SteamID + "/wishlist/")

  doc.Find(".wishlistRow").Each(func(i int, selection *goquery.Selection) {
    app := new(App)
    appID, isAppIDExist := selection.Attr("id")
    if isAppIDExist {
      app.SteamID = strings.TrimLeft(appID, "game_")
      app.Name = selection.Find(".wishlistRowItem h4.ellipsis").Text()
      app.DiscountValue = selection.Find(".discount_pct").Text()
      app.URL, _ = selection.Find(".storepage_btn_alt").Attr("href")
      p.WishList = append(p.WishList, app)
      fmt.Printf("item %d: %s\n", i, app.SteamID)
    }
  })
}
