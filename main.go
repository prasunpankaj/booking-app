package main

import (
    "fmt"
    "strings")

func main() {
    var conferenceName string = "Go Conference"
    const conferenceTickets int = 50
    var remainingTickets uint = 50
    //Fix Array
    //var bookings [50]string 

    //Dynamically add Value in Array called slice
    var bookings []string

    // type of Variable

    fmt.Printf("conferenceName type is: %T\n", conferenceName)
    fmt.Printf("conferenceTickets type is: %T\n", conferenceTickets)
    fmt.Printf("Welcome to %v booking application \n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are tickets available \n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets to attend")
    for remainingTickets > 0 && len(bookings) <= 50 {
        var firstName string
        var lastName string
        var email string
        var userTickets uint
        // Ask user to enter username

        //userName = "Pankaj Prasun"

        fmt.Printf("Enter your First Name:")
        fmt.Scan(&firstName)

        fmt.Printf("Enter your Last Name:")
        fmt.Scan(&lastName)

        fmt.Printf("Enter your Email ID:")
        fmt.Scan(&email)

        fmt.Printf("Enter how many ticket you want:")
        fmt.Scan(&userTickets)

        
        if userTickets <= remainingTickets {
            remainingTickets = remainingTickets - userTickets

        // bookings[0] = firstName + " " + lastName

            bookings = append(bookings, firstName + " " + lastName) 

            fmt.Printf("Size of the Array: %v\n", len(bookings))
            fmt.Printf("Display the Slice: %v\n", bookings)

            fmt.Printf("Thanks your %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
            
            fmt.Printf("%v ticket remaining for conference %v\n", remainingTickets, conferenceName)

            firstNames := []string{}

            for _, booking := range bookings {
                var name = strings.Fields(booking) // first name and last name into array split by array
                firstNames = append(firstNames, name[0])
            }

            fmt.Printf("The first name of booking is: %v\n", firstNames)
            fmt.Printf("These are bookings done: %v\n", bookings)

            if remainingTickets == 0 {
                fmt.Printf("All ticket has been booked. Please do visit in next year to attend conference %v\n", conferenceName)
                break
            }
        } else if userTickets == remainingTickets {

            fmt.Printf("All ticket has been booked")
         
        } else  {
            fmt.Printf("We have only %v tickets available tickets, so you can not book %v tickets\n", remainingTickets, userTickets)
            continue
        }
    }
}