package server

import (
	"log"
	"net/http"
)

func SetupServer() {
	fs := http.FileServer(http.Dir("server/temp"))

	http.Handle("/", fs)
	log.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
