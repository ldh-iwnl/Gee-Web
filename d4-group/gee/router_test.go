package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	r.addRoute("GET", "/hello/a/b/c", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/kyle")
	if n == nil {
		t.Fatal("n should not be nil")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("should be /hello/:name")
	}
	if ps["name"] != "kyle" {
		t.Fatal("name should be equal to Kyle")
	}
	fmt.Printf("matched path : %s, params['name]: %s\n", n.pattern, ps["name"])
}
func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file2.txt")
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file2.txt"
	if !ok1 {
		t.Fatal("pattern and filepath fails")
	}
	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be css/test.css")
	}
}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Print(i+1, n)
	}
	if len(nodes) != 4 {
		t.Fatal("the number of routes shoule be 4")
	}
}
