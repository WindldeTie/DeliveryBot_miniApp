package server

import "net/http"

func SetupServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "temp/index.html")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
