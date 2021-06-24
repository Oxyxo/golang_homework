package main

import (
	"fmt"
	"net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "Go is super easy!")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page text!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil) //free port on your PC after comma server parameters, sending empty - nil

}
func main() {
	handleRequest()

}

// terminal go run ... or ls - shows what is inside the folder
// go build .\main.go -> ./main.exe will run the built file
//ctrl+c -> stop the server
