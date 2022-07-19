package domain

type RouteHandler func()

type Router interface {
	Add(*RouteHandler)
	Branch(*RouteHandler)
	Root() *Router
}
