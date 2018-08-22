package base

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse struct can be used by both the base controller and other applicative controllers
// To return a
type ErrorResponse struct {
	Error string `json:"error,"`
}

// Encode encodes the struct into a JSON string
func (e ErrorResponse) Encode() string {
	msg, err := json.Marshal(e)
	if err != nil {
		panic("Failed to encode Error as JSON")
	}
	return string(msg[:])
}

// Controller is an abstract base conroller used by applicative controllers
type Controller struct {
	AllowedMethods map[string]bool
	Path           string
	Handlers       map[string]func(w http.ResponseWriter, r *http.Request)
}

// A generic match that dispatches the request according to the correlating http method
func (c *Controller) handle(w http.ResponseWriter, r *http.Request) {
	// Catch any unwanted exception
	defer func() {
		if r := recover(); r != nil {
			log.Printf("BaseHandler recovered from error %s", r)
			http.Error(w, ErrorResponse{"runtime error"}.Encode(), 500)
		}
	}()
	if methodAllowed := c.AllowedMethods[r.Method]; !(methodAllowed) {
		log.Printf("Method %s is not allowed", r.Method)
		http.Error(w, ErrorResponse{"not allowed"}.Encode(), 405)
		return
	}

	handlerFunc, exists := c.Handlers[r.Method]
	if !exists {
		log.Printf("Method %s is not implemented", r.Method)
		http.Error(w, ErrorResponse{"Method Not Implemented"}.Encode(), 501)
		return
	}
	handlerFunc(w, r)
}

// Register a controller with it's path
func (c *Controller) Register() {
	http.HandleFunc(c.Path, c.handle)
}
