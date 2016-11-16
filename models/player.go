package models

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imdoroshenko/go-steambot/loader"
)

// NewPlayer creates new Player struct.
func NewPlayer(SteamID string) *Player {
	p := new(Player)
	p.SteamID = SteamID
	return p
}

// Player struct.
type Player struct {
	SteamID  string
	WishList []*App
}

// UploadWishList uload players wishlist from players steam profile web page.
func (p *Player) UploadWishList() {
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
