package migration

import (
	"UserService/db"
	"UserService/internal/modules/user"
	"fmt"
)

func Migration() {
	db := db.Database()

	err := db.AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Users table migrated")
}
