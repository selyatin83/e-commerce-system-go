package main

import (
	"e-commerce-system/internal/infrastructure/repository/user"
	userService "e-commerce-system/internal/service/user"
	"fmt"
	"math/rand"
)

const idByGetById = 3

func main() {
	userRepository := user.CreateMemoryRepository()
	service := userService.CreateUserService(userRepository)
	addUsers(service)

	fmt.Printf("Triyng to get an user by id: %d \n", idByGetById)
	err, u := service.GetById(idByGetById)
	fmt.Printf("error: %v, user: %v \n", err, u)

	fmt.Printf("Triyng to delete an user by id: %d \n", idByGetById)
	err = service.DeleteUserById(idByGetById)
	fmt.Printf("error while deleting: %v \n", err)

	fmt.Printf("Triyng to get an user by id: %d \n", idByGetById)
	err, u = service.GetById(idByGetById)
	fmt.Printf("error: %v, user: %v \n", err, u)

	list(service)
}

func addUsers(service *userService.Service) {
	size := rand.Intn(100)
	for i := 0; i <= size; i++ {
		id := i

		err := service.CreateUser(
			id,
			fmt.Sprintf("FirstName %d", id),
			fmt.Sprintf("LastName %d", id),
			int8(rand.Intn(50)+1),
		)

		if err != nil {
			fmt.Printf("Error %s\n", err.Error())
		}
	}

	// custom user
	size++
	service.CreateUser(
		3,
		fmt.Sprintf("FirstName %d", 3),
		fmt.Sprintf("LastName %d", 3),
		int8(rand.Intn(50)+1),
	)

	fmt.Printf("%d users were added to repository \n", size)
}

func list(service *userService.Service) {
	fmt.Printf("\nUserList: \n\n")
	for _, v := range service.UserList() {
		fmt.Printf("UserId %d, Data: %v \n", v.ID, v)
	}
}
