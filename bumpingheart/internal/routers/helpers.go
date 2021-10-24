package routers

import "net/http"

type (
	DefaultHandler = func(w http.ResponseWriter, r *http.Request)
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
