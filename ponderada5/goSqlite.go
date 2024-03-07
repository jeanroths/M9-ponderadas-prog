package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
	"time"
)

func main() {
	db, _ := sql.Open("sqlite3", "./db.db")
	defer db.Close() // Defer Closing the database

	createTable(db)
	// Criando a tabela

	go Client()
	Sub(db)
	select {}
}

func createTable(db *sql.DB) {
	sqlStmt := `
  CREATE TABLE IF NOT EXISTS sensor
  (id INTEGER PRIMARY KEY, NH3_ppm INTEGER, CO_ppm INTEGER, NO2_ppm INTEGER)
  `
	// Preparando o sql statement de forma segura
	command, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Executando o comando sql
	command.Exec()
}
type Sensor struct {
	NH3_ppm, CO_ppm, NO2_ppm int
	sensor 				  	 string
}

func insertData(data Sensor, db *sql.DB) {
	// Criando uma função para inserir usuários
		_, err := db.Exec(fmt.Sprintf("INSERT INTO sensor(NH3_ppm, CO_ppm, NO2_ppm) VALUES(%d, %d, %d)", data.NH3_ppm, data.CO_ppm, data.NO2_ppm))
		if err != nil {
			log.Fatalln(err.Error())
		}
	}



func displayData(db *sql.DB) {
	row, err := db.Query("SELECT * FROM sensor ORDER BY timestamp")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var NH3_ppm int
		var CO_ppm int
		var NO2_ppm int
		var timestamp time.Time
		row.Scan(&id, &NH3_ppm, &CO_ppm, &NO2_ppm, &timestamp)
		log.Println("Sensor data: %v - %v - %v - %v - %v", id, NH3_ppm, CO_ppm, NO2_ppm, timestamp)
	}
}