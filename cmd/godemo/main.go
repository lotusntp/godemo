package main

import (
	"net/http"

	"github.com/lotusntp/godemo/pkg/app"
)

const port = ":8085"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}
