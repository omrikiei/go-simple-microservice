package monitor

import (
	"io"
	"net/http"

	"github.com/omrikiei/go-simple-microservice/base"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "up")
}

// Register registers the monitor controller to the http service
func init() {
	allowedMethods := make(map[string]bool)
	allowedMethods["GET"] = true

	handlers := make(map[string]func(w http.ResponseWriter, r *http.Request))
	handlers["GET"] = handleGet
	Monitor := base.Controller{
		AllowedMethods: allowedMethods,
		Path:           "/monitor",
		Handlers:       handlers,
	}
	Monitor.Register()
}
