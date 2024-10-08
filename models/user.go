package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

var Users = []User{
	{ID: "1", Name: "John Doe", Email: "john.doe@example.com", Phone: "123-456-7890"},
	{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com", Phone: "987-654-3210"},
}
