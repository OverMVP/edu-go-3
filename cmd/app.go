package main

import (
	"fmt"
	"log/slog"

	"l3/internal/entity"
	"l3/internal/repo"
	service "l3/internal/services"
)

func main() {
	repos := []entity.UserRepository{
		repo.NewMockUserRepo(),
		repo.NewInMemoryUserRepo(),
	}

	for _, repo := range repos {
		userService := service.NewUserService(repo)

		userService.CreateUser("Eugene", "test@gmail.com", "admin")
		userService.CreateUser("Ivan", "test2@gmail.com", "guest")

		users := userService.ListUsers()

		for _, user := range users {
			slog.Info("Users List", "Name", user.Name, "email", user.Email, "role", user.Role)
		}

		slog.Info("Looking for users by role:", "role", "admin")

		list, err := userService.FindByRole("admin")
		if err != nil {
			slog.Error("FindByRole failed", "error", err)
		}

		if len(list) > 0 {
			slog.Info("Founded users:", "count", len(list))

			for i, user := range list {
				slog.Info("User:", "number", i+1, "name", user.Name, "email", user.Email, "role", user.Role)
			}
		}

		fmt.Printf("----------------------------------------------------\n\n")
	}
}
