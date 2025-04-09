package main

import (
	"fmt"

	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/server"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	godotenv.Load()
	db.Init()
	id, _ := primitive.ObjectIDFromHex("67f56f3cff37aeb01063beba")
	user, err := db.GetUser(id)
	fmt.Println(user, err)
	server.Start()
}
