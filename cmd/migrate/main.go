package main

import (
	"ecom-tiago/configs"
	"ecom-tiago/db"
	"log"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

/*
	docs: for the migration, need to install the migrate V4
	and the migrate CLI as well.
*/
func main() {
	// docs: load the environment variables for Database Config
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User: 								configs.Envs.DBUser,
		Passwd: 							configs.Envs.DBPassword,
		Addr: 								configs.Envs.DBAddress,
		DBName: 							configs.Envs.DBName,
		Net: 									"tcp",
		AllowNativePasswords: true,
		ParseTime: 						true,
	})
	if err != nil {
		log.Fatal(err);
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err);
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err);
	}

	cmd := os.Args[(len(os.Args) - 1)];
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err);
		}
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err);
		}
	} else {
		log.Fatal("Invalid command");
	}
}