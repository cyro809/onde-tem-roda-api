package models

// TODO:  json:<field>

// User is a model
type User struct {
	ID       int    `json:"ID"`
	Nome     string `json:"Nome"`
	Email    string `json:"Email"`
	Telefone string `json:"Telefone"`
}
