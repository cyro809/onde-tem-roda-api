package models

type event struct {
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
