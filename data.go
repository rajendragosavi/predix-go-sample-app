package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func initDB(config *AppConfig) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Postgres.Hostname, config.Postgres.Port,
		config.Postgres.User, config.Postgres.Password,
		config.Postgres.DBName)

	DB, err = sql.Open("postgres", psqlInfo)
	checkErr(err)
}

func retrieveAll() []RowResult {
	statement := "SELECT type, color, location, install_date FROM playground"
	rows, err := DB.Query(statement)
	checkErr(err)

	items := []RowResult{}
	for rows.Next() {
		var rowResult = RowResult{}
		err := rows.Scan(
			&rowResult.EquipType,
			&rowResult.Color,
			&rowResult.Location,
			&rowResult.InstallDate)
		checkErr(err)
		items = append(items, rowResult)
	}
	return items
}
