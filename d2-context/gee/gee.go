package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New constructor
func New() *Engine {
	return &Engine{router: newRouter()}
}

// add route to router
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c :=newContext(w,req)
	engine.router.handle(c)
}
