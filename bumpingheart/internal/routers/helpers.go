package routers

import "net/http"

type (
	DefaultHandler = func(ctx *Context) error
	WrappedRouter  interface {
		Setup(router DefaultHandler, handlers ...DefaultHandler)
	}
)

type HttpMethods string

func (h HttpMethods) String() string {
	return string(h)
}

const (
	MethodGet    HttpMethods = "GET"
	MethodPost   HttpMethods = "POST"
	MethodPut    HttpMethods = "PUT"
	MethodDelete HttpMethods = "DELETE"
	MethodPatch  HttpMethods = "PATCH"
)

// RegisterContextHandler convert a DefaultHandler to a http.HandlerFunc,
// by creating and passing a Context to it.
func RegisterContextHandler(handler DefaultHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{
			request: r,
			writer:  w,
		}
		handler(ctx)
	}
}
