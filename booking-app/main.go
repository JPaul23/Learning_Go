package main

import (
	"fmt"
	"sync"
	"time"
)

// package level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// arrays & slice
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// wait group for goroutines
var wg = sync.WaitGroup{}

func main() {

	//functions
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	//check user tickets
	if isValidName && isValidEmail && isValidTicketNumber {

		// booking function
		bookTicket(userTickets, firstName, lastName, email)

		// go routines
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		//IF ELSE
		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address does not contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets is invalid.")
		}
	}

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here")
}

func getFirstNames() []string {
	//slice
	firstNames := []string{}
	//looping in list
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user inputs
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Println("Enter your Email address: ")
	fmt.Scan(&email)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v, you have %v tickets, You will recieve a confirmation email at %v\n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// go routines, concurrency
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}
