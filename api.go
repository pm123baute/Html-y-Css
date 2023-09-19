package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
)

type Parameters struct {
	User  string
	Email string
	Pass  string
}

var params Parameters

type login struct {
	User  string `json:"user"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		panic(err)
	}
	var lo login
	err = json.Unmarshal(body, &lo)
	if err != nil {
		panic(err)
	}
	user := lo.User
	email := lo.Email
	pass := lo.Pass
	fmt.Printf("User: %s, Email: %s, Pass: %s\n", user, email, pass)
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Configura los manejadores de CORS para permitir solicitudes desde todos los orígenes.
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	// Crea el enrutador de HTTP y registra el manejador de rutas "/api".
	http.Handle("/api", cors(http.HandlerFunc(handleRequests)))

	// Inicia el servidor en el puerto 8080.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
