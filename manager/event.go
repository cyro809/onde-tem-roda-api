package manager

import (
	"fmt"

	"github.com/onde-tem-roda-api/database"
	"github.com/onde-tem-roda-api/models"
)

// funções para o consultas e alterações no banco de dados
// Veirifcar função In e Select (projeto do cartola, manger amigos.go)

// SelectEventByID selects event from the database by its ID
func SelectEventByID(ID int) models.Event {
	event := models.Event{}
	db, err := database.OpenDB()
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM event WHERE ID = %d", ID))
	checkErr(err)
	defer db.Close()

	// Verificar o uso de .Get ou Select
	for rows.Next() {
		rows.Scan(
			&event.ID,
			&event.NomeEvento,
			&event.Grupo,
			&event.Responsavel,
			&event.ResponsavelTel,
			&event.ResponsavelEmail,
			&event.Endereco,
			&event.PlaceID,
			&event.UserID,
		)
	}

	return event
}

// SelectEventsByField selects events by field value likeness
func SelectEventsByField(field string, value string) []models.Event {
	var events []models.Event

	db, err := database.OpenDB()
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM event WHERE %s LIKE "%%%s%%"`, field, value))
	checkErr(err)
	defer db.Close()

	// Verificar o uso de .Get ou Select
	for rows.Next() {
		event := models.Event{}
		rows.Scan(
			&event.ID,
			&event.NomeEvento,
			&event.Grupo,
			&event.Responsavel,
			&event.ResponsavelTel,
			&event.ResponsavelEmail,
			&event.Endereco,
			&event.PlaceID,
			&event.UserID,
		)
		events = append(events, event)
	}

	return events
}

// SelectAllEvents selects all events
func SelectAllEvents() []models.Event {
	var events []models.Event

	db, err := database.OpenDB()
	checkErr(err)

	rows, err := db.Query("SELECT * FROM event")
	checkErr(err)
	defer db.Close()

	// Verificar o uso de .Get ou Select
	for rows.Next() {
		event := models.Event{}
		rows.Scan(
			&event.ID,
			&event.NomeEvento,
			&event.Grupo,
			&event.Responsavel,
			&event.ResponsavelTel,
			&event.ResponsavelEmail,
			&event.Endereco,
			&event.PlaceID,
			&event.UserID,
		)
		events = append(events, event)
	}

	return events
}
