package models

// TODO:  json:<field>

// Event struct to be used in the API and the database
type Event struct {
	ID               int     `json:"ID"`
	NomeEvento       string  `json:"NomeEvento"`
	Grupo            string  `json:"Grupo"`
	Responsavel      string  `json:"Responsavel"`
	ResponsavelTel   string  `json:"ResponsavelTel"`
	ResponsavelEmail string  `json:"ResponsavelEmail"`
	Endereco         string  `json:"Endereco"`
	PlaceID          string  `json:"PlaceID"`
	Latitude	     float32 `json:"Latitude"`
	Longitude		 float32 `json:"Longitude"`
	UserID           int     `json:"UserID"`
}
