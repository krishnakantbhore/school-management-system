package model

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      string `json:"age"`
}

type Teacher struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Subject   string `json:"subject"`
	Class     string `json:"class"`
	Email     string `json:"email"`
}