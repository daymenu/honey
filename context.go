package honey

import (
	"net/http"
	"sync"
)

// Context 定义
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	index    int8
	fullPath string

	engine *Engine
	mu     sync.RWMutex

	keys map[string]interface{}

	sameSite http.SameSite
}

func (c *Context) reset() {
	c.Writer = &c.writermem
	c.index = -1
}

func (c *Context) String(body string) {
	c.writermem.WriteString(body)
}
