package domain

type User struct {
	Name     string
	Did      string
	Email    string
	Password string
	Post     []*post
}
