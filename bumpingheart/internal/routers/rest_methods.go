package routers

import (
	"net/http"
)

type RESTMethods struct {
	List   DefaultHandler
	Create DefaultHandler
	Update DefaultHandler
	Delete DefaultHandler
}

// Create a net.HandlerFunc for each method
func (methods RESTMethods) RegisterMethod(writer http.ResponseWriter, request *http.Request) {
	// Create the context object
	context := &Context{
		writer:  writer,
		request: request,
	}
	context.initLocals()

	// Check if the method is GET
	if request.Method == MethodGet.String() {
		// Check the List method is not nil
		if methods.List != nil {
			methods.List(context)
		}
	} else if request.Method == MethodPost.String() {
		// Check the Create method is not nil
		if methods.Create != nil {
			methods.Create(context)
		}
	} else if request.Method == MethodPut.String() {
		// Check the Update method is not nil
		if methods.Update != nil {
			methods.Update(context)
		}
	} else if request.Method == MethodDelete.String() {
		// Check the Delete method is not nil
		if methods.Delete != nil {
			methods.Delete(context)
		}
	}
}
