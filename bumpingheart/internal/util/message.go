package util

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type WrappedMessage struct {
	Err            error
	Message        string
	HttpStatusCode uint16
	Type           string
	IssuedAt       time.Time
}

func WriteJSON(w http.ResponseWriter, message WrappedMessage) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(message.HttpStatusCode))

	// Convert the message to JSON
	messageBytes, err := json.Marshal(&message)
	if err == nil {
		w.Write(messageBytes)
	} else {
		w.Write([]byte(`{"code": 500,"error": "Failed to marshal message"}`))
	}

	return err
}

func GetDynamicRoute(request *http.Request) string {
	path := strings.SplitAfter(request.URL.Path, "/")
	if len(path) > 1 {
		return path[len(path)-1]
	}
	return ""
}
