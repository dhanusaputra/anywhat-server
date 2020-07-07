package main

import (
	"database/sql"
	"fmt"

	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	name     = "postgres"
	password = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := service.NewAnywhatService(db)
	if err := cmd.ListenGRPC(s, "8080"); err != nil {
		panic(err)
	}
}
