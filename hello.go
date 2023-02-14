package main

import "fmt"

func main() {
	// const organizer = "Usman Khalil"
	// var conferenceName = "Go Conference"
	// const conferenceTickets = 50
	// var remainingTickets = 50

	// fmt.Printf("conferenceTickets is %T type variable \n", conferenceTickets)

	// fmt.Println("Welcome to our booking application conference called", conferenceName, ". Be sure to enjoy")
	// fmt.Printf("%v is responsible for managing all sort of stuff \n", organizer)
	// fmt.Println("We have total of", conferenceTickets, "tickets and", "only", remainingTickets, "seats are available")

	// // static typing
	// var userName string
	// var userTickets int
	// // Asking user about his user name
	// userName = "Evan"
	// userTickets = 1
	// fmt.Printf("%v brought %v tickets", userName, userTickets)

	// // infering type
	// guestOfHonour := "M. Azeez"
	// fmt.Printf("Our guest of honor is %v", guestOfHonour)

	// var guestUser string

	// fmt.Print("Do you want to enter the conference as well? Add your name here")
	// fmt.Scan(&guestUser)

	// fmt.Printf("%v thank you. Your name has been registered", guestUser)

	// Lets do some mathematics

	// var firstNumber int
	// var secondNumber int

	// fmt.Println("Hey, you wanna do some summation? Give me two numbers and I can add them for you")
	// fmt.Println("Enter first number ")
	// fmt.Scan(&firstNumber)

	// fmt.Println("Great, now the second one ")
	// fmt.Scan(&secondNumber)

	// var total = firstNumber + secondNumber
	// fmt.Printf("These two numbers sum up to %v", total)

	// Arrays
	// var friends = [40]string{"Dan", "Rob", "Evan"}

	// Array without initialization
	// var frameworks [3]string

	// frameworks[0] = "React"
	// frameworks[1] = "Angular"
	// frameworks[2] = "Vue"

	// fmt.Printf("My top %v favorite frameworks are %v", len(frameworks), frameworks)

	// // Slices: dynamic array

	// var friends []string

	// friends = append(friends, "Qasim")

	// Todos List
	var todos []string

	for {
		var todo string = ""

		fmt.Println("Add a new task")

		fmt.Scan(&todo)

		todos = append(todos, todo)
		todo = ""

		fmt.Printf("Here's your current tasks list %v", todos)
	}

}
