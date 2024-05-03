package main

import (
	"database/sql"
	"log"

	"github.com/atiwat-r/golang-backend/cmd/api"
	"github.com/atiwat-r/golang-backend/config"
	"github.com/atiwat-r/golang-backend/db"
	"github.com/go-sql-driver/mysql"
)





func main() {

	db_cfg := mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}
	db, err := db.NewMySQLStorage(db_cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	err = server.Run()

	if (err != nil) {
		log.Fatal(err)
	}
}




func initStorage(db *sql.DB) {
	err := db.Ping()
	if (err != nil) {
		log.Fatal(err)
	}

	log.Println("DB: Successfully Connected")
}