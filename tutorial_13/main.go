package main

import "fmt"

type Player struct {
	Name   string
	Health int
}

// Value Receiver: Does NOT change the original health
func (p Player) CheckStatus() {
	fmt.Printf("%s has %d HP remaining.\n", p.Name, p.Health)
}

// Pointer Receiver: DOES change the original health
func (p *Player) TakeDamage(amount int) {
	p.Health -= amount
	fmt.Printf("%s took %d damage!\n", p.Name, amount)
}

func main() {
	hero := Player{Name: "GopherKnight", Health: 100}

	hero.CheckStatus()
	hero.TakeDamage(20) // Original health is now 80
	hero.CheckStatus()
}