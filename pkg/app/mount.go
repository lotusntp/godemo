package app

import "net/http"

//Mount handler to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index)
	mux.HandleFunc("news", newsView)

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/edit", adminEdit)
	// ทุก path ที่อยู่ใน adminMux จะต้อง middleware onlyAdmin
	mux.Handle("/admin/", onlyAdmin(adminMux))
	// /news/:path
	// /admin/login
	// /admin/post
	// /admin/update

}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
