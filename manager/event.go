package manager

import (
	"fmt"

	"github.com/onde-tem-roda-api/database"
	"github.com/onde-tem-roda-api/models"
)

// funções para o consultas e alterações no banco de dados
// Veirifcar função In e Select (projeto do cartola, manger amigos.go)

// SelectEventByID selects user from the database by its ID
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
		fmt.Println("next")
	}

	return event
}
