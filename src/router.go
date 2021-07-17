package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

//SetRouter 設定路由
func SetRouter(r *mux.Router) error {
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`<font color="red">Hello red world!</font>`))
	})
	/*
		r.HandleFunc("/products", ProductsHandler)
	*/
	return nil
}
