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
		for opsi != 8 {
			if opsi == 1 {
				eventt.AddEvents(&event, &n)

			} else if opsi == 2 {
				eventt.ShowEvents(event, n)
			} else if opsi == 3 {
				eventt.UpdateEvents(&event, n)
			} else if opsi == 4 {
				eventt.SearchEvent(event, n)
			} else if opsi == 5 {
				eventt.DeleteEvent(&event, &n)
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
	fmt.Println("4) Search Event")
	fmt.Println("5) Delete Eventt")
	fmt.Println("6) lorem ipsum dolor sit amet, consectetur adipiscing elit")
	fmt.Println("8) Exit")
	fmt.Println()
	fmt.Println("Enter Option")

	var opsi int
	fmt.Scan(&opsi)
	return opsi
}
