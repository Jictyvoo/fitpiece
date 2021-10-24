package bumpingheart

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jictyvoo/fitpiece/bumpingheart/internal/routers"
	"github.com/jictyvoo/fitpiece/bumpingheart/internal/util"
)

type RESTRouter struct {
	Namespace        string
	BatchHandlers    routers.RESTMethods
	SpecificHandlers routers.RESTMethods
	Middlewares      []routers.DefaultHandler
	SubRouters       routers.WrappedRouter
	Parameters       string
}

func (rest RESTRouter) setupBatch(restRouter *http.ServeMux) {
	// Register List method using HTTP.get
	defaultNamespace := "/" + rest.Namespace
	if rest.BatchHandlers.List != nil {
		if len(rest.Parameters) > 0 {
			restRouter.Handle(defaultNamespace+"/"+rest.Parameters, routers.RegisterContextHandler(rest.BatchHandlers.List))
		}
	}
}

func (rest RESTRouter) forbiddenMethod(writer http.ResponseWriter, request *http.Request) {
	message := util.WrappedMessage{
		Err:            fmt.Errorf("%d: Method not allowed", http.StatusMethodNotAllowed),
		Message:        "Can't create `" + rest.Namespace + "` directly with an ID",
		HttpStatusCode: http.StatusMethodNotAllowed,
		Type:           "CLIENT_REQUEST",
		IssuedAt:       time.Now(),
	}

	_ = util.WriteJSON(writer, message)
}

func (rest RESTRouter) registerHandlers(writer http.ResponseWriter, request *http.Request) {
	// Check if request path is empty
	if len(request.URL.Path) == len(rest.Namespace)+1 && request.URL.Path == rest.Namespace+"/" {
		// Register Batch methods in the router with defaultNamespace
		rest.BatchHandlers.RegisterMethod(writer, request)
	} else {
		log.Println("Request path: ", request.URL.Path)
		if request.Method == routers.MethodPost.String() {
			// Check the Create method is not nil
			rest.forbiddenMethod(writer, request)
		} else {
			rest.SpecificHandlers.RegisterMethod(writer, request)
		}
	}
}

func (rest RESTRouter) Setup(router *http.ServeMux, handlers ...routers.DefaultHandler) {
	rest.setupBatch(router)

	// Register handlers in the router
	router.Handle(rest.Namespace+"/", http.HandlerFunc(rest.registerHandlers))

	// Create a sub router only for routes that have an ID
	/*if rest.SubRouters != nil {
		groupWithID := restGroup.Group("/:id")
		rest.SubRouters.Setup(groupWithID)
	}*/
}
