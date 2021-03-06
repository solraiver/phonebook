package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/solraiver/phonebook/entyti"
	"log"
	"net/http"
)

type bookHandler struct {
	srv entyti.PhoneBookService
}

func NewBookHandler(s entyti.PhoneBookService) *bookHandler {
	return &bookHandler{srv: s}
}

//var autoincrement int
//autoincrement = 0

func (h bookHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Лог1")
	id := ps.ByName("id")
	log.Println("ЭТО БЛЯДЬ АЙДИ", id)
	user := h.srv.GetUser(id)

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		log.Println(err.Error())
	}
	log.Println("Пользователь", user)
}

/*func (h bookHandler) getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := Users
	if len(user) == 0 {
		log.Println("Пользователи не найдены!")
		return
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
	}
	log.Println("Список пользователей!")
}

func (h bookHandler) addUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}
	key := strconv.Itoa(autoincrement)

	if user.Firstname != "" {
		Users[key] = user
		log.Println("Пользователь", Users[key].Firstname, "добавлен!")
	}
	autoincrement++
}

func (h bookHandler) updateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}
	if _, ok := Users[id]; !ok {
		log.Println("Пользователь не найден!")
	} else {
		Users[id] = user
		log.Println("Пользователь обновлен!")
	}
}

func (h bookHandler) deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")
	_, ok := Users[id]
	if !ok {
		log.Println("Пользователь не найден!")
		return
	} else {
		delete(Users, id)
		log.Println("Пользователь удален!")
	}

}
*/
