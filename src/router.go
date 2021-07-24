package router

import (
	"io/ioutil"
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
	r.HandleFunc("/post", postHandler).Methods(`POST`)
	return nil
}
func homeHandler(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte(`<font color="red">Hello red world!</font>`))
	http.ServeFile(w, req, path.Join(`res/index.html`))
}
func postHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		// Will return the post query which function got.
		w.Write(body)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	//w.Write([]byte(`Request Method: ` + req.Method))
}
