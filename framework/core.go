package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

func (c *Core) Get(pattern string, handler ControllerHandler) {
	c.router[pattern] = handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, req *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(req, response)

	router := c.router["fool"]
	if router == nil {
		return
	}
	log.Println("core.router")
	router(ctx)
}
