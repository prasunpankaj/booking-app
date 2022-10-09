package main

import (
    "fmt"
    "booking-app/helper"
    )

    //Package Level Variables
    const conferenceTickets uint = 50
    var remainingTickets uint = 50
    var conferenceName  = "Go Conference"
    var bookings = make([]UserData, 0) //List of Map and Intization with size 0

    type UserData struct {
        firstName string
        lastName string
        email string
        numberOfTickets uint
    }

func main() {

    greetUser()

    fmt.Printf("We have total of %v tickets and %v are tickets available \n", conferenceTickets, remainingTickets)
   
    fmt.Println("Get your tickets to attend")
    for remainingTickets > 0 && len(bookings) <= 50 {
        
        firstName, lastName, email, userTickets := getUserInput()
        
        fmt.Println("remainingTickets is in Main %v", remainingTickets)

        isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email ,userTickets, remainingTickets)

        if  isValidName && isValidEmail && isValidTicket {
            
            bookTicket(userTickets , firstName , lastName, email)

            fmt.Printf("Thanks your %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
            
            fmt.Printf("%v ticket remaining for conference %v\n", remainingTickets, conferenceName)

            firstNames := getFirstName()

            fmt.Printf("The first name of booking is: %v\n", firstNames)

            if remainingTickets == 0 {
                fmt.Printf("All ticket has been booked. Please do visit in next year to attend conference %v\n", conferenceName)
                break
            }
        } else if userTickets == remainingTickets {

            fmt.Printf("All ticket has been booked")
         
        } else  {

            if !isValidName {
                fmt.Println("Please enter valid first name and last name and min two characters")
            }

            if !isValidEmail {
                fmt.Println("Please enter valid email address")
            }

            if !isValidTicket {
                fmt.Println("Please enter valid Ticket number, We have only %v tickets available tickets, so you can not book %v tickets\n", remainingTickets, userTickets)
            }

            fmt.Println("Your input data is invaild, please try again!")
            //fmt.Printf("We have only %v tickets available tickets, so you can not book %v tickets\n", remainingTickets, userTickets)
            continue
        }
    }
} // End of Main function

func greetUser() {
    fmt.Printf("conferenceName type is: %T\n", conferenceName)
    fmt.Printf("conferenceTickets type is: %T\n", conferenceTickets)
    fmt.Printf("Welcome to %v booking application \n", conferenceName)
} // End of greetUser


func getFirstName() []string {
    firstNames := []string{}

    for _, booking := range bookings {
        //var name = strings.Fields(booking) // first name and last name into array split by array
        firstNames = append(firstNames, booking.firstName)
    }

    fmt.Printf("These are bookings done: %v\n", bookings)

    return firstNames
}

func getUserInput() (string,string,string,uint) {
    var firstName string
    var lastName string
    var email string
    var userTickets uint

    fmt.Printf("Enter your First Name:")
    fmt.Scan(&firstName)

    fmt.Printf("Enter your Last Name:")
    fmt.Scan(&lastName)

    fmt.Printf("Enter your Email ID:")
    fmt.Scan(&email)

    fmt.Printf("Enter how many ticket you want:")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string)  {

    fmt.Printf("RemainingTickets: %d\n", remainingTickets)

    remainingTickets = remainingTickets - userTickets

    var userData = UserData{
        firstName: firstName,
        lastName: lastName,
        email: email,
        numberOfTickets: userTickets,
    }

    bookings = append(bookings, userData)
    fmt.Printf("Size of the Array: %v\n", len(bookings))
    fmt.Printf("Display the Slice: %v\n", bookings)
}