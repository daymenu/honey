package honey

import (
	"net/http"
	"sync"
)

// HandlerFunc 定义处理http的方法签名
type HandlerFunc func(*Context)

// HandlersChain 定义一个HandlerFunc的数组
type HandlersChain []HandlerFunc

// Last 返回最后一个HandlerFunc
func (c HandlersChain) Last() HandlerFunc {
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
}

// RouteInfo 路由信息
type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc HandlerFunc
}

// RouteInfos 路由切片
type RouteInfos []RouteInfo

// Engine 定义web项目引擎
type Engine struct {
	trees   methodTrees
	context *Context
	pool    sync.Pool
}

// New 返回Engine实例
func New() *Engine {
	engine := &Engine{}
	engine.pool.New = func() interface{} {
		return engine.allocateContext()
	}
	return engine
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()
	engine.handleHTTPRequest(c)
	engine.pool.Put(c)
}

func (engine *Engine) allocateContext() *Context {
	return &Context{engine: engine}
}

func (engine *Engine) addRoute(method, path string, handlers HandlersChain) {
	root := engine.trees.get(method)
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		engine.trees = append(engine.trees, methodTree{method: method, root: root})
	}

	root.addRoute(path, handlers)
}
func (engine *Engine) handleHTTPRequest(c *Context) {

}
