package main

import (
	"Q_two/database"
	"Q_two/handle"
)

func main() {
	db := database.Connect("identifier.sqlite")
	server := handle.NewServer(db)
	server.RUN()

}
