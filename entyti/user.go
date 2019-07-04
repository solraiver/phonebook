package entyti

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
}

type PhoneBookService interface {
	GetUser(id string) User
}
