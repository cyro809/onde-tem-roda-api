package database

// AddEventDateColumn checks if column exists otherwise adds column to event table
func AddEventDateColumn() {
	db, err := OpenDB()
	var userVersion int
	st1 := "ALTER TABLE event ADD COLUMN EventDate TEXT"
	st2 := "PRAGMA user_version = 1"
	statements := []string{st1, st2}

	rows, err := db.Query("PRAGMA user_version")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for rows.Next() {
		rows.Scan(&userVersion)
	}

	if userVersion == 0 {
		for _, st := range statements {
			statement, _ := db.Prepare(st)
			statement.Exec()
		}
	}
}
