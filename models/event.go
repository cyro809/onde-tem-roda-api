package models

// TODO:  json:<field>

// Event struct to be used in the API and the database
type Event struct {
	ID               int
	NomeEvento       string
	Grupo            string
	Responsavel      string
	ResponsavelTel   string
	ResponsavelEmail string
	Endereco         string
	PlaceID          string
	UserID           int
}
