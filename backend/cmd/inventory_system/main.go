package main

import (
	"fmt"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/db"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	//"github.com/LostProgrammer1010/InventorySystem/internal/server"
	"github.com/LostProgrammer1010/InventorySystem/internal/server/api/post"
	"github.com/joho/godotenv"
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

	err = db.AddUser(test)

	user := models.User{Username: "LostProgrammer", Password: "123"}

	jwtToken, err := post.LoginInUser(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jwtToken)
	//server.Start()
}
