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

	rows, err := db.Query(fmt.Sprintf("SELECT ID, NomeEvento, Grupo, Responsavel, ResponsavelTel, ResponsavelEmail, Endereco, PlaceID, Latitude, Longitude, UserID, EventDate FROM event WHERE ID = %d", ID))
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
			&event.Latitude,
			&event.Longitude,
			&event.UserID,
			&event.EventDate,
		)
	}

	return event
}

// SelectEventsByField selects events by field value likeness
func SelectEventsByField(field string, value string) []models.Event {
	var events []models.Event

	db, err := database.OpenDB()
	checkErr(err)

	rows, err := db.Query(fmt.Sprintf(`SELECT ID, NomeEvento, Grupo, Responsavel, ResponsavelTel, ResponsavelEmail, Endereco, PlaceID, Latitude, Longitude, UserID, EventDate FROM event WHERE %s LIKE "%%%s%%"`, field, value))
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
			&event.Latitude,
			&event.Longitude,
			&event.UserID,
			&event.EventDate,
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

	rows, err := db.Query("SELECT ID, NomeEvento, Grupo, Responsavel, ResponsavelTel, ResponsavelEmail, Endereco, PlaceID, Latitude, Longitude, UserID, EventDate FROM event")
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
			&event.Latitude,
			&event.Longitude,
			&event.UserID,
			&event.EventDate,
		)
		events = append(events, event)
	}

	return events
}

// InsertEvent inserts event on database
func InsertEvent(event models.Event) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare("INSERT INTO event (NomeEvento, Grupo, Responsavel, ResponsavelTel, ResponsavelEmail, Endereco, PlaceId, Latitude, Longitude, UserID, EventDate) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")

	statement.Exec(
		event.NomeEvento,
		event.Grupo,
		event.Responsavel,
		event.ResponsavelTel,
		event.ResponsavelEmail,
		event.Endereco,
		event.PlaceID,
		event.Latitude,
		event.Longitude,
		event.UserID,
		event.EventDate,
	)
}

// UpdateEventByField updates user name
func UpdateEventByField(ID int, field string, value string) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare(fmt.Sprintf(`UPDATE event SET %s=? WHERE id = ?`, field))
	checkErr(err)

	res, err := statement.Exec(value, ID)
	checkErr(err)

	affec, err := res.RowsAffected()
	checkErr(err)

	checkAffectedRows(affec)
}

// DeleteEventByID deletes event from database by its ID
func DeleteEventByID(ID int) {
	db, err := database.OpenDB()
	checkErr(err)

	statement, err := db.Prepare("DELETE FROM event WHERE id = ?")
	checkErr(err)

	res, err := statement.Exec(ID)
	checkErr(err)

	affec, err := res.RowsAffected()
	checkErr(err)

	checkAffectedRows(affec)
}
