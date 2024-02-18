package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type fooHandler struct{}
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (foo *fooHandler) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	user := new(User)
	err := json.NewDecoder(reader.Body).Decode(user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, err)
		return
	}

	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, string(data))
}

func barHandler(writer http.ResponseWriter, reader *http.Request) {
	name := reader.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(writer, "Hello %s!", name)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, reader *http.Request) {
		fmt.Fprint(writer, "Hello Golang!")
	})

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)
}
