package Gym

import "fmt"

type Member struct {
	ID     int
	Name   string
	Active bool
}

type Gym struct {
	Name    string
	Members []Member
}

// Constructor
func NewGym(name string) *Gym {
	return &Gym{
		Name:    name,
		Members: []Member{},
	}
}

// Add member
func (g *Gym) AddMember(member Member) {
	g.Members = append(g.Members, member)
}

// Activate membership
func (g *Gym) ActivateMember(id int) {
	for i, m := range g.Members {
		if m.ID == id {
			g.Members[i].Active = true
		}
	}
}

// Print members
func (g Gym) PrintMembers() {
	fmt.Println("Gym members:")
	for _, m := range g.Members {
		fmt.Printf("ID: %d | Name: %s | Active: %v\n", m.ID, m.Name, m.Active)
	}
}
