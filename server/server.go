package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func SetupServer() {
	fs := http.FileServer(http.Dir("server/temp"))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Убедимся, что путь не пытается выйти за пределы temp
		path := filepath.Join("server/temp", r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// Если файл не найден — пытаемся отдать index.html
			http.ServeFile(w, r, "server/temp/index.html")
			return
		}
		// Иначе — отдаём обычный файл
		fs.ServeHTTP(w, r)
	}))
	
	log.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
