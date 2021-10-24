package routers

import (
	"net/http"
	"time"

	"github.com/jictyvoo/fitpiece/bumpingheart/internal/util"
)

type Context struct {
	request *http.Request
	writer  http.ResponseWriter
	locals  map[string]interface{}
}

func (ctx Context) GetRouteParam() string {
	return util.GetDynamicRoute(ctx.request)
}

func (ctx Context) Request() *http.Request {
	return ctx.request
}

func (ctx Context) Send(data []byte) error {
	if _, err := ctx.writer.Write(data); err != nil {
		return err
	}
	return nil
}

func (ctx Context) SendString(data string) error {
	if _, err := ctx.writer.Write([]byte(data)); err != nil {
		return err
	}
	return nil
}

func (ctx Context) SendWrapped(data util.WrappedMessage) error {
	if err := util.WriteJSON(ctx.writer, data); err != nil {
		return err
	}
	return nil
}

func (ctx Context) SendError(err error, customStatus ...int) error {
	status := http.StatusInternalServerError
	if len(customStatus) > 0 {
		status = customStatus[0]
	}

	// Write the status code
	ctx.writer.WriteHeader(status)
	writeErr := util.WriteJSON(
		ctx.writer,
		util.WrappedMessage{
			Err: err, HttpStatusCode: uint16(status), IssuedAt: time.Now(),
		})
	if writeErr != nil {
		return writeErr
	}
	return nil
}

func (ctx *Context) initLocals() {
	// Initialize the map if it's nil
	if ctx.locals == nil {
		ctx.locals = make(map[string]interface{})
	}
}

// Set sets a key value pair in the context
func (ctx *Context) Set(key string, value interface{}) {
	// Initialize the map if it's nil
	ctx.initLocals()
	ctx.locals[key] = value
}

// Get gets a value from the context
func (ctx Context) Get(key string) interface{} {
	if ctx.locals != nil {
		return ctx.locals[key]
	}
	return nil
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.writer.Header().Set(key, value)
}

func (ctx Context) GetHeader(key string) string {
	return ctx.request.Header.Get(key)
}
