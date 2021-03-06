package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/solraiver/phonebook/handlers"
	"github.com/solraiver/phonebook/store"
	"log"
	"net/http"
)

func main() {

	books := store.NewPhoneBook()
	hr := handlers.NewBookHandler(books)

	router := httprouter.New()
	router.GET("/user/:id", hr.GetUser)

	/*router.GET("/user", getUsers)
	router.POST("/user", addUser)
	router.DELETE("/user/:id", deleteUser)
	router.PUT("/user/:id", updateUser)*/

	log.Println("Я запускаюсь")
	log.Fatal(http.ListenAndServe(":8080", router))

}
