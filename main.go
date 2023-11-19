package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"runtime"

	_ "github.com/lib/pq"
	"github.com/lordfarshad/simplebank/api"
	db "github.com/lordfarshad/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:@localhost:5432/simplebank?sslmode=disable"
	serverAddress = "localhost:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	openURL(serverAddress)
}

func openURL(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to open URL: %v\n", err)
	}
}
