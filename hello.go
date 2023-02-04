package main

import (
	"fmt"
)

func main() {
	const organizer = "Usman Khalil"
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T type variable \n", conferenceTickets)

	fmt.Println("Welcome to our booking application conference called", conferenceName, ". Be sure to enjoy")
	fmt.Printf("%v is responsible for managing all sort of stuff \n", organizer)
	fmt.Println("We have total of", conferenceTickets, "tickets and", "only", remainingTickets, "seats are available")

	// static typing
	var userName string
	var userTickets int
	// Asking user about his user name
	userName = "Evan"
	userTickets = 1
	fmt.Printf("%v brought %v tickets", userName, userTickets)

	// infering type
	guestOfHonour := "M. Azeez"
	fmt.Printf("Our guest of honor is %v", guestOfHonour)
	fmt.Print("Nice")

}
