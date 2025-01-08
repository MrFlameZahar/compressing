package main

import (
	ports "bd/bd/internal/ports/http"
	repo "bd/bd/internal/repo/redis"
	"net/http"
)

func main() {
	go repo.InitRedis()
	http.ListenAndServe(":8050", ports.NewRouter())
}
