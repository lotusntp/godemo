package app

import (
	"net/http"

	"github.com/lotusntp/godemo/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
