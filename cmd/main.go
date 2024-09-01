package main

import (
	"database/sql"
	"ecom-tiago/cmd/api"
	"ecom-tiago/configs"
	"ecom-tiago/db"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// docs: load the environment variables for Database Config
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// docs: Check if the database is connected
	initStorage(db)

	// docs: Run a new API Server
	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected!")
}
