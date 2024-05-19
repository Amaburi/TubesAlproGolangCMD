package utils

import (
	"fmt"
	"math/rand"
)

type details struct {
	EventId                                 int
	Name, Date, Location, Description, Time string
}

const maxx = 20

type Events [maxx]details

func AddEvents(isEvents *Events, n *int) bool {

	var continueAdding string
	for {
		if *n > maxx {
			*n = maxx
			return false
		}

		fmt.Print("Enter Event Name:")
		fmt.Scan(&isEvents[*n].Name)

		fmt.Print("Enter Event Date:")
		fmt.Scan(&isEvents[*n].Date)

		fmt.Print("Enter Event Time:")
		fmt.Scan(&isEvents[*n].Time)

		fmt.Print("Enter Event Location:")
		fmt.Scan(&isEvents[*n].Location)

		fmt.Print("Enter Event Description:")
		fmt.Scan(&isEvents[*n].Description)

		isEvents[*n].EventId = rand.Intn(190)

		*n++
		fmt.Println("Do you want to continue adding events? (yes/no)")
		fmt.Scan(&continueAdding)

		if continueAdding == "no" {
			return false
		}

	}
}

func ShowEvents(e Events, n int) bool {
	var continueAdding string
	for i := 0; i < n; i++ {
		fmt.Println("Id:", e[i].EventId)
		fmt.Println("Event Name:", e[i].Name)
		fmt.Println("Event Date:", e[i].Date)
		fmt.Println("Event Time:", e[i].Time)
		fmt.Println("Event Location:", e[i].Location)
		fmt.Println("Event Description:", e[i].Description)
		fmt.Println()

	}
	fmt.Println("Do you want to back? (yes/no)")
	fmt.Scan(&continueAdding)

	return continueAdding == "yes"
}

func UpdateEvents(ev *Events, n int) {
	var option, options string
	var num int
	var found bool

	fmt.Println("Which attributes would you like to change?")
	fmt.Scan(&option)
	fmt.Println("input the desired value")
	fmt.Scan(&options)
	fmt.Println("input the id of the event")
	fmt.Scan(&num)

	for i := 0; i < n; i++ {
		if option == "Name" && ev[i].EventId == num {
			ev[i].Name = options

		} else if option == "Date" && ev[i].EventId == num {
			ev[i].Date = options
		} else if option == "Time" && ev[i].EventId == num {
			ev[i].Time = options
		} else if option == "Location" && ev[i].EventId == num {
			ev[i].Location = options
		} else if option == "Description" && ev[i].EventId == num {
			ev[i].Description = options
		}
		found = ev[i].EventId == num
	}
	if found {
		fmt.Print("Event has been updated successfully\n")
	} else {
		fmt.Print("The id you inputted is not in the array\n")
	}
}
func DeleteEvent(ev *Events, n *int) {
	var num, left, right, mid, idx int

	fmt.Println("input the id of the event")
	fmt.Scan(&num)
	left = 0
	right = *n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if ev[mid].EventId == num {
			ev[mid] = ev[mid+1]
			*n = mid + 1
			fmt.Print("Data has been deleted")
		} else if ev[mid].EventId < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if ev[mid].EventId != num {
			fmt.Println("The id you're looking for is not in the array")
		}

	}
}
func SearchEvent(ev Events, n int) bool {
	var num, left, right, mid, idx int
	var continueSearching string
	fmt.Println("input the id of the event")
	fmt.Scan(&num)
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if ev[mid].EventId == num {
			fmt.Println("Id:", ev[mid].EventId)
			fmt.Println("Event Name:", ev[mid].Name)
			fmt.Println("Event Date:", ev[mid].Date)
			fmt.Println("Event Time:", ev[mid].Time)
			fmt.Println("Event Location:", ev[mid].Location)
			fmt.Println("Event Description:", ev[mid].Description)
			idx = mid

		} else if ev[mid].EventId < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
		if ev[mid].EventId != num {
			fmt.Println("The id you're looking for is not in the array")
		}
		fmt.Println("Do you want to continue Searching events? (yes/no)")
		fmt.Scan(&continueSearching)

		if continueSearching == "no" {
			return false
		}

	}

	return true
}
