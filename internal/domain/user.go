package domain

// User defines the properties of a user
type User struct {
	Id       int64
	UserName string
	Email    string
}

// NewUser creates a new user with the given parameters
func NewUser(id int64, userName, email string) *User {
	return &User{
		Id:       id,
		UserName: userName,
		Email:    email,
	}
}
