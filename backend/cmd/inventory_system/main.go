package main

import (
	"fmt"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"github.com/LostProgrammer1010/InventorySystem/internal/server"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	godotenv.Load()
	db.Init()
	test := models.User{
		Username:     "LostProgrammer",
		Password:     "123",
		Role:         "admin",
		Organization: []models.Organization{{Name: "NSVC", Role: "admin"}},
	}
	token, err := authentication.CreateRefreshToken(test, "test")
	if err != nil {
		fmt.Println(err)
	}

	test.RefreshToken = append(test.RefreshToken, *token)

	db.CreateUser(test)

	id, _ := primitive.ObjectIDFromHex("67f69d23a00b21b765fedb6d")
	fmt.Println(db.GetUser(id))
	server.Start()
}
