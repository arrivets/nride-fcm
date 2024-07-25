package store

type User struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type IStore interface {
	AddUser(id string, token string) User
	DeleteUser(id string)
	GetUser(id string) *User
}
