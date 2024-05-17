package utils

import "fmt"

type Users struct {
	Username string
	Password string
}

func Register(userr *Users) bool {

	var pass string
	fmt.Print("Enter username: ")
	fmt.Scan(&userr.Username)

	fmt.Print("Enter password: ")
	fmt.Scan(&userr.Password)

	for i := 0; i < 5; i++ {
		fmt.Print("Renter password: ")
		fmt.Scan(&pass)
		if userr.Password == pass {
			fmt.Println("Registration successful!")
			return true
		} else {
			fmt.Println("Passwords do not match. Please try again. Maximum attempts are 5")
		}
	}
	fmt.Println("Registration failed. Maximum attempts reached.")
	return false
}
