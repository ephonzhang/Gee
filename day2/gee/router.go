package gee

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Req.Method + "-" + c.Req.URL.Path

	if hander, ok := r.handlers[key]; ok {
		hander(c)
	} else {
		c.String(c.StatusCode, "404 NOT FOUND, URL=%s", c.Req.URL)
	}
}
