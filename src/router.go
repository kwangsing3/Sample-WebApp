package router

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

var flagMode = int8(0)

//SetRouter 設定路由
func SetRouter(r *mux.Router, flag int8) error {
	flagMode = flag
	r.HandleFunc("/", homeHandler)
	/*
		r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(`<font color="red">Hello red world!</font>`))
		})
	*/
	return nil
}
func homeHandler(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte(`<font color="red">Hello red world!</font>`))
	http.ServeFile(w, req, path.Join(`res/index.html`))
}
