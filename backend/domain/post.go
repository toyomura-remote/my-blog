package domain

type post struct {
	did     string
	title   string
	content string
	user    *User
}
