package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", cfg.FormatDSN()) // "root:mypassword@tcp(localhost:3306)/db"
	if (err != nil) {
		log.Fatal(err)
	}

	return db, nil
}

/** Docker MySQL server

// Create Container/MySQL server "golang-db"
docker run --name golang-db -e MYSQL_ROOT_PASSWORD=mypassword -p 3306:3306 mysql

// Enter the container shell
docker exec -it golang-db bash

// Login to MySQL as root
mysql -u root -p

// SQL statement to show all Databases in the MySQL server
SHOW DATABASES;

// Create new Database "db"
CREATE DATABASE db;

*/



