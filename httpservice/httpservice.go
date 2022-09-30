package httpservice

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler() {
	r := mux.NewRouter()

	r.Methods("GET").Path("/getBalance/{id:[0-9]+}").HandlerFunc(getBalance)
	r.Methods("POST").Path("/addToBalance/{id:[0-9]+}").HandlerFunc(addToBalance)
	r.Methods("POST").Path("/deleteToBalance/{id:[0-9]+}").HandlerFunc(delBalance)
	r.Methods("POST").Path("/transactionToUser/{idSend:[0-9]+}/{idRecipe:[0-9]+}").HandlerFunc(exchangeBetweenUsers)
	http.Handle("/", r)
}
