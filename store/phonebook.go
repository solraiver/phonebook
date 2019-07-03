package store

import "github.com/solraiver/phonebook/entyti"

type phoneBook struct {
	list map[string]entyti.User
}

func NewPhoneBook() *phoneBook {
	return &phoneBook{list: map[string]entyti.User{}}
}

func (p phoneBook) GetUser() {

}
