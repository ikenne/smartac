package main

import (
	"deviceAPI/api"
	"deviceAPI/database"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	port = "8086"
)

func main() {

	db := database.NewTestDB()
	apiServer := api.NewServer(port, db)

	fmt.Println("Starting API server ...")
	apiServer.Start()

	defer func() {
		fmt.Println("Stopping API server.")
		apiServer.Stop()
		fmt.Println("Stopped.")
	}()

	terminated := make(chan os.Signal)
	signal.Notify(terminated, syscall.SIGINT, syscall.SIGTERM)
	<-terminated
}
