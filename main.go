package main

import (
	"encoding/json"
	"go-package/config"
	"go-package/controller"
	"go-package/middleware"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/users", controller.Index).Methods("GET")

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		// CARA 1
		// json.NewEncoder(w).Encode(map[string]bool{"ok": true})

		// Cara 2
		res, _ := json.Marshal(map[string]bool{"ok": true})
		w.Write(res)
	}).Methods("GET")

	r.Use(middleware.LoggingMiddleware)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
