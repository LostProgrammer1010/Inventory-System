package main

import (
	"context"

	"github.com/LostProgrammer1010/Inventory-System/internal/server"
	"github.com/LostProgrammer1010/Inventory-System/internal/server/db"
)

func main() {

	client := db.Connect()

	defer client.Disconnect(context.TODO())

	server.Start()
}
