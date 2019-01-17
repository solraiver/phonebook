package main

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type User struct {
	Id          string
	Firstname   string
	Lastname    string
	Phonenumber string
}

var Users = make([]User, 0)

func main() {
	///u := User{Id: "1", Firstname: "And", Lastname: "Kuku", Phonenumber: "111"}
	///Users = append(Users, u)

	Users = []User{
		{Id: "1", Firstname: "And", Lastname: "Kuku", Phonenumber: "111"},
		{Id: "vasya", Firstname: "And", Lastname: "Kuku", Phonenumber: "111"},
	}

	router := httprouter.New()
	router.GET("/user/:id", getUser)
	log.Println("Я запускаюсь")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func findUser(users []User, id string) (User, error) {
	for _, v := range users {
		if v.Id == id {
			return v, nil
		}
	}
	return User{}, errors.New("данного пользователя не существует")
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")

	user, err := findUser(Users, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
	}
	//1
}
