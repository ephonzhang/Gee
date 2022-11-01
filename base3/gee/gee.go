package gee

import (
	"fmt"
	"net/http"
)

type HanderFunc func (http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HanderFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HanderFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, hander HanderFunc) {
	key := method + "-" + pattern
	engine.router[key] = hander
}

func (engine *Engine) GET (pattern string, handler HanderFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST (pattern string, handler HanderFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP (w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path

	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}