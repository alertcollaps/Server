package httpservice

import (
	"Server/payment"
	"Server/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type answer struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello!")
	fmt.Fprintf(w, "Hello!")
}

func GetLast(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello!")
	last, err := sql.GetLastElement()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Fprintf(w, last)
}

func GetAlpha(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "httpservice/templates/Index.html")
}

func SayLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	var err error
	var message int
	defer func() {
		sendMessage(w, err, "Balance: "+strconv.Itoa(message))
	}()
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	message, err = payment.GetBalance(idInt)
}

func exchangeBetweenUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		sendMessage(w, err, "Exchange between users: ")
	}()
	params := mux.Vars(r)
	idSend := params["idSend"]
	idRecipe := params["idRecipe"]
	idSendInt, err := strconv.Atoi(idSend)
	if err != nil {
		return
	}
	idRecipeInt, err := strconv.Atoi(idRecipe)
	if err != nil {
		return
	}

	p := sql.Balance{}
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		err = fmt.Errorf("can't parse json")
		return
	}

	err = payment.ExchangeBetweenUsers(idSendInt, idRecipeInt, p.Cash)
}

func addToBalance(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		sendMessage(w, err, "Adding to balance")
	}()
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	p := sql.Balance{}
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		err = fmt.Errorf("can't parse json")
		return
	}

	err = payment.AddSum(idInt, p.Cash)
	if err != nil {
		return
	}
}

func delBalance(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		sendMessage(w, err, "Decrease balance")
	}()
	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("can't parse to int")
		return
	}
	p := sql.Balance{}
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		err = fmt.Errorf("can't parse json")
		return
	}

	err = payment.SubtractionSum(idInt, p.Cash)
	if err != nil {
		return
	}
	return
}

func sendMessage(w http.ResponseWriter, err error, message string) {
	answer := answer{}
	answer.Error = "None"
	answer.Message = message
	if err != nil {
		answer.Error = err.Error()
	}

	b, err := json.Marshal(answer)
	if err != nil {
		return
	}

	_, err = w.Write(b)
	if err != nil {
		return
	}
}
