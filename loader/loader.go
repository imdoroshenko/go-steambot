package loader

import (
  "net/http"
  "fmt"
  "github.com/PuerkitoBio/goquery"
)

const attrProp string = "attr"

func Get (url string) (*goquery.Document, error) {
  resp, err := http.Get(url)
  defer resp.Body.Close()

  if err != nil {
    fmt.Println("http get error")
    return nil, err;
  }

  doc, err := goquery.NewDocumentFromResponse(resp)

  if err != nil {
    fmt.Println("goquery error")
    return nil, err;
  }
  //result := make(map[string][]string)
  //for name, query := range queries {
  //  result[name] = ExtractFromDocument(doc, query[0], query[1], query[2])
  //}
  return doc, nil

  //doc.Find(query).Each(func(i int, s *goquery.Selection) {
  //  var value string
  //  switch prop {
  //    case attrProp:
  //      value, _ = s.Attr(name)
  //    break
  //    default:
  //      value = s.Text()
  //  }
  //  result = append(result, value)
  //  fmt.Printf("item %d: %s\n", i, value)
  //})
  //return result, nil
}

func ExtractFromDocument (doc *goquery.Document, query, prop, name string) []string {
  var result []string
  doc.Find(query).Each(func(i int, s *goquery.Selection) {
    var value string
    switch prop {
    case attrProp:
      value, _ = s.Attr(name)
      break
    default:
      value = s.Text()
    }
    result = append(result, value)
    fmt.Printf("item %d: %s\n", i, value)
  })
  return result
}
