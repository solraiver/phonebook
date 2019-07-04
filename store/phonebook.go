package store

import "github.com/solraiver/phonebook/entyti"

type phoneBook struct {
	list map[string]entyti.User
}

func NewPhoneBook() *phoneBook {
	return &phoneBook{list: map[string]entyti.User{}}
}

func (p *phoneBook) GetUser(id string) entyti.User {
	p.list["0"] = entyti.User{FirstName: "Ivan", LastName: "Ivanov", PhoneNumber: "02"}

	return p.list[id]
}
