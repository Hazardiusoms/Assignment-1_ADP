package main

import (
	"Assignment1/Employee"
	"Assignment1/Gym"
	"Assignment1/Hotel"
	"Assignment1/Wallet"
	"fmt"
)

func main() {

	// HOTEL
	hotel := Hotel.NewHotel("Astana Plaza")
	hotel.AddRoom(Hotel.Room{Number: 101, Type: "Single", PricePerNight: 50})
	hotel.BookRoom(101)
	hotel.PrintRooms()

	fmt.Println()

	// EMPLOYEES
	staff := Employee.Staff{}
	staff.AddEmployee(Employee.Employee{
		ID: 1, Name: "Nurlybek", Position: "Manager", BaseSalary: 2000,
	})
	staff.PrintSalaries()

	fmt.Println()

	// GYM
	gym := Gym.NewGym("Power Gym")
	gym.AddMember(Gym.Member{ID: 1, Name: "Aruzhan"})
	gym.ActivateMember(1)
	gym.PrintMembers()

	fmt.Println()

	// WALLET
	wallet := Wallet.Wallet{}
	wallet.Deposit(500)

	err := wallet.Pay(120)
	if err != nil {
		fmt.Println("Payment failed:", err)
	} else {
		fmt.Println("Payment successful. Balance:", wallet.Balance)
	}
}
