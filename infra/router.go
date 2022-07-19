package infra

import (
	"fmt"
	"go-terminal/domain"
)

type route struct {
	handler domain.RouteHandler
	before  *route
	routes  map[string]*route
}

type Router struct {
	route
}

var pos *route

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Routing(scan string) {
	var ok bool
	pos, ok = r.routes[scan]
	if !ok {
		fmt.Printf("no routing: %s\n", scan)
		return
	}

	pos.handler()
}
