package main

import (
	"fmt"
	"golang-api/internal/api"
	"golang-api/internal/models"
	"golang-api/internal/service"
	"golang-api/internal/storage"
)

func main() {
	st := storage.New()
	srv := service.New(st)
	handlers := api.New(srv)

	Server(handlers)
}

func Server(handlers api.IApi) {
	handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_2",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	user, _ := handlers.GetUser(&models.GetUserRequest{
		1,
	})
	fmt.Printf("%v", user)

	handlers.UpdateUser(&models.UpdateUserRequest{
		Id:       int(user.Id),
		Login:    "updated_manager",
		Password: "12345678",
		Name:     "Daler_",
		Surname:  "Sultonov_",
	})

	handlers.DeleteUser(&models.DeleteUserRequest{
		Id: 1,
	})

	handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	handlers.AddUser(&models.AddUserRequest{
		Login:    "manager_1",
		Password: "123456",
		Name:     "Daler",
		Surname:  "Sultonov",
	})

	users := handlers.GetUsers(&models.GetUsersRequest{})
	fmt.Printf("%v", users)

	statistics := handlers.GetStatustics()
	fmt.Println(statistics)
}
