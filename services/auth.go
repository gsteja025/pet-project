package services

type User struct {
	Username string
	Email    string
}

func NewUser(user_name string, Email string) (*User, error) {

	user := &User{
		Username: user_name,
		Email:    Email,
	}

	return user, nil

}
