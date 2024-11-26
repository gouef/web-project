package models

type User struct {
	ID    int
	Name  string
	Email string
}

// Mock databáze
var users = []User{
	{ID: 1, Name: "Jan", Email: "jan@example.com"},
	{ID: 2, Name: "Petr", Email: "petr@example.com"},
}

func GetAllUsers() []User {
	return users
}

func GetUserByID(id int) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
