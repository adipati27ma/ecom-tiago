package main

import (
	"database/sql"
	"ecom-tiago/cmd/api"
	"ecom-tiago/config"
	"ecom-tiago/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main()  {
	// docs: load the environment variables for Database Config
	db, err := db.NewMySQLStorage(mysql.Config{
		User: 								config.Envs.DBUser,
		Passwd: 							config.Envs.DBPassword,
		Addr: 								config.Envs.DBAddress,
		DBName: 							config.Envs.DBName,
		Net: 									"tcp",
		AllowNativePasswords: true,
		ParseTime: 						true,
	})
	if err != nil {
		log.Fatal(err);
	}

	// docs: Check if the database is connected
	initStorage(db);
	
	// docs: Run a new API Server
	server := api.NewAPIServer(":8080", db);
	if err:= server.Run(); err != nil {
		log.Fatal(err);
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping();
	if err != nil {
		log.Fatal(err);
	}

	log.Println("Database connected!");
}