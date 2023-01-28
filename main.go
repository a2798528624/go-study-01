package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("www"))
	http.Handle("/", fileServer)

	http.HandleFunc("/login", loginHandler)

	fmt.Println("Listening on port 8081", "http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Println(err)
	}
	username := request.FormValue("username")
	password := request.FormValue("password")
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}
