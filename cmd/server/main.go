package main

import (
	"database/sql"
	"fmt"

	"github.com/dhanusaputra/anywhat-server/pkg/cmd"
	"github.com/dhanusaputra/anywhat-server/pkg/service"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 12345
	user     = "admin"
	password = "123456"
	dbname   = "anywhat"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
