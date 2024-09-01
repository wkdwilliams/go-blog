package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/wkdwilliams/go-blog/internal/infrastructure"
)

func main() {
	command := flag.String("command", "", "Migrate up or down")
	flag.Parse()

	if *command != "up" && *command != "down" {
		fmt.Println("Usage: migrate -command [up,down]")
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := infrastructure.NewMysql()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(conn, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if *command == "up" {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}
