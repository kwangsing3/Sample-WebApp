package main

import (
	router "WebApp/src"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var port = ":1337"
var Flag = int8(0)
var AcceptFlags = []string{`-dev`, `-pro`} // 0:Debug Mode、1:Production mode

func main() {

	// Configure launch mode
	argnames := os.Args
	msg := ``
	if len(argnames) == 2 {
		for i := range AcceptFlags {
			Flag = map[bool]int8{true: int8(i), false: 0}[strings.EqualFold(argnames[1], AcceptFlags[i])]
		}
	} else {
		for i := range AcceptFlags {
			msg += `[` + AcceptFlags[i] + `] `
		}
		fmt.Println(`This application can run with one args tag. Option: ` + msg)
	}

	msg = `Start to run with flag:` + AcceptFlags[Flag]
	fmt.Println(msg)

	// Setting Router
	r := mux.NewRouter()
	router.SetRouter(r, Flag) //Setting website Path

	fmt.Println("Server Listen at: " + port)
	http.ListenAndServe(port, r)
	return
}
