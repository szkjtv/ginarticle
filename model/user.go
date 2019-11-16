package model

type User struct {
	Id       int    `json:"id,omitempty"`
	Account  string `json:"account,omitempty"`
	Password string `json:"password,omitempty"`
}
