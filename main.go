package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type User struct {
	Firstname   string
	Lastname    string
	Phonenumber string
}

var Users = make(map[string]User)
var autoincrement int

func main() {
	autoincrement = 0
	Users["0"] = User{Firstname: "Ivan", Lastname: "Ivanov", Phonenumber: "02"}

	router := httprouter.New()
	router.GET("/user/:id", getUser)
	router.POST("/user", addUser)
	router.DELETE("/user/:id", deleteUser)
	log.Println("Я запускаюсь")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")

	user, ok := Users[id]
	if !ok {
		log.Println("Пользователь не найден!")
		return
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
	}
}

func addUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}

	key := string(autoincrement)
	Users[key] = user

	autoincrement++
}

func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")
	_, ok := Users[id]
	if !ok {
		log.Println("Пользователь не найден!")
		return
	}

	delete(Users, id)

}
