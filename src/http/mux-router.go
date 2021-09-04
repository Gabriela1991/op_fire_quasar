package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	if uri != "/" {
		muxDispatcher.HandleFunc(uri, f).Methods("GET")
	} else {
		muxDispatcher.HandleFunc(uri, f).Methods("GET")
	}
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	if len(port) != 0 {
		fmt.Println("Server route mux running in port :" + port)
	} else {
		port = "8000"
		fmt.Println("Server route mux running in port DEFAULT:" + port)
	}
	http.ListenAndServe(":"+port, muxDispatcher)
}

func (*muxRouter) SWAGGERUI(handler http.Handler) {
	muxDispatcher.PathPrefix("/swaggerui").Handler(http.StripPrefix("/swaggerui", handler))
}
