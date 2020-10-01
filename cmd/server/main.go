package server

import (
	"log"
	"net/http"
)

func Run(port string, assetDir string) {
	err := http.ListenAndServe(port, http.FileServer(http.Dir(assetDir)))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
