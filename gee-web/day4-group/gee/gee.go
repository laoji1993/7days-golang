package gee

type HandlerFunc func(*Context)

type (
	RouterGroup struct {
		prefix      string
		middlewares []*HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}
	Engine struct {
		*RouterGroup
		router router
	}
)
