package routers

import (
	"net/http"
)

type GroupMethods struct {
	Get    DefaultHandler
	Post   DefaultHandler
	Put    DefaultHandler
	Patch  DefaultHandler
	Delete DefaultHandler
}

// Create a net.HandlerFunc for each method
func (methods GroupMethods) RegisterMethod(writer http.ResponseWriter, request *http.Request) {
	// Create the context object
	context := &Context{
		writer:  writer,
		request: request,
	}
	context.initLocals()

	// Check if the method is GET
	if request.Method == MethodGet.String() {
		// Check the Get method is not nil
		if methods.Get != nil {
			methods.Get(context)
		}
	} else if request.Method == MethodPost.String() {
		// Check the Post method is not nil
		if methods.Post != nil {
			methods.Post(context)
		}
	} else if request.Method == MethodPut.String() {
		// Check the Put method is not nil
		if methods.Put != nil {
			methods.Put(context)
		}
	} else if request.Method == MethodPatch.String() {
		// Check the Patch method is not nil
		if methods.Patch != nil {
			methods.Patch(context)
		}
	} else if request.Method == MethodDelete.String() {
		// Check the Delete method is not nil
		if methods.Delete != nil {
			methods.Delete(context)
		}
	}
}
