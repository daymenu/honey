package honey

import (
	"bytes"
)

const (
	paramsWildcard   byte = ':'
	catchAllWildcard byte = '*'
	regexWildcard    byte = '{'
)

var endRegexWildcard []byte = []byte("}")

const (
	pathSplitChar byte = '/'
)

// Param 通配符参数
type Param struct {
	Key   string
	Value string
}

// Params 存储参数的切片
type Params []Param

// Get 获取通配符参数
func (params Params) Get(name string) (string, bool) {
	for _, entry := range params {
		if entry.Key == name {
			return entry.Value, true
		}
	}
	return "", false
}

// ByName 根据名字获取通配符参数
func (params Params) ByName(name string) (va string) {
	va, _ = params.Get(name)
	return
}

type methodTree struct {
	method string
	root   *node
}

type methodTrees []methodTree

func (trees methodTrees) get(method string) *node {
	for _, tree := range trees {
		if tree.method == method {
			return tree.root
		}
	}
	return nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

func countParams(path string) uint16 {
	var n uint
	for i := range []byte(path) {
		switch path[i] {
		case paramsWildcard, catchAllWildcard, regexWildcard:
			n++
		}
	}
	return uint16(n)
}

type nodeType uint8

const (
	static nodeType = iota //default
	root
	param    // :
	catchAll // *
	regex    //{}
)

type node struct {
	path      string        // 路径
	nType     nodeType      //树节点类型
	priority  uint32        // 优先级
	children  []*node       // 存储孩子节点
	handlers  HandlersChain // 该节点关联的方法
	fullPath  string        // 全路径
	wildChild bool          // 是否是通配符
	regex     string        // 解析出来的正则表达式
	paramName string        // 解析出的参数名字
}

func (n *node) addRoute(path string, handlers HandlersChain) {
	fullPath := path
	n.priority++
	if len(n.path) == 0 && len(n.children) == 0 {
		n.insertChild(path, fullPath, handlers)
		n.nType = root
		return
	}
}

// findWildcard 查找通配符
func findWildcard(path string) (wildcard string, i int, valid bool) {
	for start, c := range []byte(path) {
		if c != paramsWildcard && c != catchAllWildcard && c != regexWildcard {
			continue
		}
		valid = true
		for end, c := range []byte(path[start+1:]) {
			switch c {
			case pathSplitChar:
				return path[start : start+1+end], start, valid
			case paramsWildcard, catchAllWildcard, regexWildcard:
				valid = false
			}
		}
		return path[start:], start, valid
	}
	return "", -1, false
}

func parseWildcard(wildcard string) (nType nodeType, paramName string, wildcardValue string) {
	wildcardByte := []byte(wildcard)
	if wildcardByte[0] == paramsWildcard {
		return param, wildcard[1:], string(paramsWildcard)
	}

	if wildcardByte[0] == catchAllWildcard {
		return catchAll, wildcard[1:], string(catchAllWildcard)
	}

	if wildcardByte[0] == regexWildcard {
		i := bytes.LastIndex(wildcardByte, endRegexWildcard)
		return regex, wildcard[i+1:], wildcard[1:i]
	}
	return
}

func (n *node) insertChild(path string, fullPath string, handlers HandlersChain) {
	n.path = path
	n.handlers = handlers
	n.fullPath = fullPath
}
