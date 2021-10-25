package main

import (
	"log"
	"net/http"

	bumpingheart "github.com/jictyvoo/fitpiece/bumpingheart"
)

// Handler function that check GET request and returns a list of strings
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// Create new http server
	serveMux := http.NewServeMux()
	genericRestRouter := bumpingheart.RESTRouter{
		Namespace: "/api",
		BatchHandlers: bumpingheart.RESTMethods{
			List: func(ctx *bumpingheart.Context) error {
				return ctx.SendString("List")
			},
		},
		SpecificHandlers: bumpingheart.RESTMethods{
			List: func(ctx *bumpingheart.Context) error {
				id := ctx.GetRouteParam()
				return ctx.Send([]byte("Specific List " + id))
			},
		},
	}

	genericRestRouter.Setup(serveMux)
	serveMux.HandleFunc("/", handler)

	// Start the server, listening on port 8080
	log.Println("Server Started!\nListening on port 8080")
	if err := http.ListenAndServe(":8080", serveMux); err != nil {
		log.Fatalln(err)
	}
}
