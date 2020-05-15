package router

import "net/http"

// Context 上下文结构体
type Context struct {
	request *http.Request
}
