package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	FirstName, LastName string
	Age                 uint16
	Balance             int16
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("Hi i am %s %s. I am %d years old", u.FirstName, u.LastName, u.Age)
}

func (u *User) setName(name string) {
	u.FirstName = name
}

func homePage(w http.ResponseWriter, r *http.Request) {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
		Balance:   1000,
	}
	tmpl, _ := template.ParseFiles("templates/homepage.html")
	tmpl.Execute(w, user)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
