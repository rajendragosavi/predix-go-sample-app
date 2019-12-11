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

	stmt, err1 := DB.Prepare(" CREATE TABLE IF NOT EXISTS playground (equip_id serial PRIMARY KEY,type varchar (50) NOT NULL,color varchar (25) NOT NULL,location varchar(25) check (location in ('north', 'south', 'west', 'east', 'northeast', 'southeast', 'southwest', 'northwest')),install_date date);")
	checkErr(err1)
	_, err1 = stmt.Exec()
	checkErr(err1)

	stmt, err1 = DB.Prepare("INSERT INTO public.playground(   type, color, location, install_date) VALUES ('pump', 'blue', 'south', '2017-11-03');")
	checkErr(err1)
	_, err1 = stmt.Exec()
	checkErr(err1)
	
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
