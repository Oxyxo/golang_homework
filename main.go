package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16 //full number over 0
	money                 int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
} // *User and User difference
func home_page(page http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	// bob.name = "Alex" //will change name to Alex
	// fmt.Fprintf(page, "User name is:"+bob.name)
	bob.setNewName("Alex")
	fmt.Fprintf(page, bob.getAllInfo()) //User name is: Bob. He is 25 and he has money equal: -50
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
	//var bob User = ...
	// bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}

	handleRequest()

}

// terminal go run ... or ls - shows what is inside the folder
// go build .\main.go -> ./main.exe will run the built file
//ctrl+c -> stop the server
