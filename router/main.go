package router

import (
  "net/http"
  "fmt"
)

type Router struct {
  Actions []*Action
}

func (r *Router) AddActions(actions ...*Action) *Router {
  for _, action := range actions {
    r.Actions = append(r.Actions, action.Compile())
  }
  return r;
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
  fmt.Println("ServeHTTP")
  var result []byte

  for _, action := range r.Actions {
    if action.IsMatch(req) {
      result = action.Handler(res, req, action.ExtractParams(req))
      break
    }
  }

  if result == nil {
      res.Write([]byte(""))
    } else {
      res.Write(result)
  }
}
