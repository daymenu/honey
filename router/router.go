package router

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义请求执行体
type HandlerFunc func(*Context)

// Engine router 结构体
type Engine struct {
	router map[string]http.HandlerFunc
}

func (engine *Engine) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		fmt.Fprintf(responseWriter, "URL.Path = %q\n", request.URL.Path)
	case "/hello":
		for k, v := range request.Header {
			fmt.Fprintf(responseWriter, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(responseWriter, "404 NOT FOUND: %s\n", request.URL)
	}
}

// New 新建router
func New() *Engine {
	return new(Engine)
}

func (engine *Engine) addRoute(method string, parttern string, handle http.HandlerFunc) {

}
