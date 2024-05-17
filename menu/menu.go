package menu

import (
	"fmt"
	eventt "tubes/utils"
)

func Run() {
	var user eventt.Users
	var event eventt.Events
	var opsi, n int

	register := eventt.Register(&user)

	if register {
		fmt.Println("Hello,", user.Username)
		opsi = Menu()
		for opsi != 6 {
			if opsi == 1 {
				eventt.AddEvents(&event, &n)

			} else if opsi == 2 {
				eventt.ShowEvents(event, n)
			} else if opsi == 3 {
				eventt.UpdateEvents(&event, n)
			}
			fmt.Println()
			fmt.Println("Hello,", user.Username)
			opsi = Menu()
		}
	} else {
		fmt.Print("user needs to be registered")
	}
}

func Menu() int {
	fmt.Println("==================")
	fmt.Println("       Menu")
	fmt.Println("==================")

	fmt.Println("1) Add Event")
	fmt.Println("2) Show All Event")
	fmt.Println("3) Update Event")
	fmt.Println("4) Add Event")
	fmt.Println("5) Add Event")
	fmt.Println("6) Exit")
	fmt.Println()
	fmt.Println("Enter Option")

	var opsi int
	fmt.Scan(&opsi)
	return opsi
}
