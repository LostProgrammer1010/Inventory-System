package main

import (
	"github.com/LostProgrammer1010/Inventory-System/internal/db"
	"github.com/LostProgrammer1010/Inventory-System/internal/server"
)

func main() {

	db.Connect()

	server.Start()
}
