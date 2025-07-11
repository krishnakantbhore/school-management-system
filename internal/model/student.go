package model

type Student struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Class     string `json:"class"`
	Email     string `json:"email"`
}