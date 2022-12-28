package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]User, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//for {
	// Ask user for details
	firstName, lastName, email, requiredTickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, requiredTickets)

	if isValidName && isValidEmail && isValidTickets {
		bookTickets(firstName, lastName, email, requiredTickets)
		wg.Add(1)
		go sendTickets(firstName, lastName, email, requiredTickets)
		// Print only first names
		printFirstNames()
	} else {
		if !isValidName {
			fmt.Println("firt name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTickets {
			fmt.Println("number of tickets you entered is invalid")
		}
		//continue
	}

	//}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The First names of bookings are %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var requiredTickets uint
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets to be booked: ")
	fmt.Scan(&requiredTickets)

	return firstName, lastName, email, requiredTickets
}

func bookTickets(firstName string, lastName string, email string, requiredTickets uint) {
	remainingTickets = remainingTickets - requiredTickets

	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: requiredTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("Thanks you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, requiredTickets, email)
	fmt.Printf("There are %v tickets remaining for the %v\n", remainingTickets, conferenceName)
	fmt.Printf("Total bookings are %v\n", bookings)
}

func sendTickets(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################################################################")
	fmt.Printf("Sending tickets: %v sent to email %v\n", ticket, email)
	fmt.Println("####################################################################")
	wg.Done()
}
