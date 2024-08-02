package main

import (
	"fmt"
	"gostuff/helper"
	"time"
)

var conferenceName string = "GoInfo"
var remainingTickets int = 50

const conferenceTickets int = 50

var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {
	getUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isUserTicketsNumber := helper.ValidInformation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isUserTicketsNumber {
			bookTickets(userTickets, firstName, lastName, email)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := printFirstName()
			fmt.Printf("First names in booking %v\n", firstNames)

			noTicketsBooking := remainingTickets == 0
			if noTicketsBooking {
				fmt.Printf("All conference tickets are booked.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("It's not a valid name")
			}
			if !isValidEmail {
				fmt.Println("It's not a valid email")
			}
			if !isUserTicketsNumber {
				fmt.Println("It's not a valid ticket number")
			}
			fmt.Printf("Your input data is invalid. Please enter new values.\n")
		}
	}
}

func getUsers() {
	fmt.Printf("My Conference Name is %v\n", conferenceName)
	fmt.Printf("Number of seats %v count, %v available.\n", remainingTickets, conferenceTickets)
}

func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Print("Please enter your first name:\n ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name:\n ")
	fmt.Scan(&lastName)
	fmt.Print("Please enter your email: \n")
	fmt.Scan(&email)
	fmt.Print("Please enter your user tickets: \n")
	fmt.Scan(&userTickets)
	fmt.Printf("My first name is: %v\n", firstName)
	fmt.Printf("My last name is: %v\n", lastName)
	fmt.Printf("My Email address is: %v\n", email)
	fmt.Printf("User tickets info: %v\n", userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)
	fmt.Printf("Remaining tickets in conference: %v\n", remainingTickets)
	//fmt.Printf("Please print your booking: %v\n", bookings)
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket: %v  to email address to: %v \n", tickets, email)
	fmt.Println("#########")
}
