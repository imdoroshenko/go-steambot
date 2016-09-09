package router

import (
  "net/http"
  "regexp"
  "strings"
)

type ActionHandler func(http.ResponseWriter, *http.Request, map[string]string) []byte

const rePrefix string = "^"
const rePostfix string = "$"
const reParamMatch string = `\[(\w+)]`
const reParamPlaceholder string = "(.+)"

func NewAction (method string, path string, handler ActionHandler) *Action {
  action := new(Action)
  action.Method = method
  action.Path = path
  action.Handler = handler
  return action
}

type Action struct {
  Method string
  Path string
  Handler ActionHandler
  re *regexp.Regexp
  paramsList []string
}

func (a Action) IsMatch(req *http.Request) bool {
  return a.re.MatchString(req.URL.Path) && req.Method == a.Method
}

func (a Action) ExtractParams(req *http.Request) map[string]string {
  params := make(map[string]string)
  paramsMatches := a.re.FindAllStringSubmatch(req.URL.Path, -1)

  for i, match := range(paramsMatches[0][1:]) {
    params[a.paramsList[i]] = match
  }
  return params;
}

func (a *Action) Compile() *Action {
  patters := regexp.MustCompile(reParamMatch)
  matches := patters.FindAllStringSubmatch(a.Path, -1)
  rePattern := rePrefix + a.Path + rePostfix
  for _, match := range(matches) {
    a.paramsList = append(a.paramsList, match[1])
    rePattern = strings.Replace(rePattern, match[0], reParamPlaceholder, -1)
  }
  a.re = regexp.MustCompile(rePattern)
  return a;
}
