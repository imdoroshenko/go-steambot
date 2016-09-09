package models
//
//import "github.com/imdoroshenko/go-steambot/loader"
//
//var detailsQuery = map[string][]string{
//  "name": []string{"apphub_AppName", "text", nil},
//  "discount": []string{"..game_area_purchase_game_wrapper .game_purchase_discount .discount_pct", "text", nil}}
//
//func newApp (steamID string) *App {
//  a := new(App)
//  a.SteamID = steamID
//  return a
//}
//
type App struct {
  SteamID string
  Name string
  DiscountValue string
  URL string
}
//
//func (a *App) UploadDetails () *App {
//  loader.Get("http://store.steampowered.com/app/291190/")
//  return a
//}
