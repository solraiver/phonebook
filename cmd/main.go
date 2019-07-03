package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/solraiver/phonebook/store"
	"log"
	"net/http"
)





func main() {

	phoneBook := store.NewPhoneBook()
	phoneBook.

	//Users["0"] = User{Firstname: "Ivan", Lastname: "Ivanov", Phonenumber: "02"}

	router := httprouter.New()
	router.GET("/user/:id", getUser)
	router.GET("/user", getUsers)
	router.POST("/user", addUser)
	router.DELETE("/user/:id", deleteUser)
	router.PUT("/user/:id", updateUser)

	log.Println("Я запускаюсь")
	log.Fatal(http.ListenAndServe(":8080", router))

}


