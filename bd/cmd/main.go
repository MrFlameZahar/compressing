package main

import (
	ports "bd/bd/internal/ports/http"
	"net/http"
)

func main() {

	http.ListenAndServe(":8050", ports.NewRouter())
}
