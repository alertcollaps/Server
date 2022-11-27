package httpservice

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").HandlerFunc(GetAlpha)
	r.Methods("GET").Path("/hello").HandlerFunc(SayHello)

	http.Handle("/", r)
}
