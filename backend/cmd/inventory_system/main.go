package main

import (
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.Init()
	server.Start()
}
