package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type VendingMachine struct {
	water, milk, beans, cash, cups int
}

type CupOfCoffee struct {
	water, milk, beans, cup, price int
	name                           string
	idx                            int
}

func (v *VendingMachine) Buy(m map[int]CupOfCoffee) {
	var id, money int
	defer DisplayGreeting()
	p := fmt.Println
	pf := fmt.Printf

	for key, val := range m {
		pf("%d:  %s\n", key, val.name)
	}

	fmt.Fscan(os.Stdin, &id)
	if id > len(m) {
		fmt.Println("Choose valid index type of coffee")
		return
	}

	ChosenCoffeeType := m[id]
	pf("You choose %s, price is %d$, please insert money...", ChosenCoffeeType.name, ChosenCoffeeType.price)
	fmt.Fscan(os.Stdin, &money)

	if money < ChosenCoffeeType.price {
		p("Not enough money (((")
		return
	}

	canMakeCupOfCoffee := false
	switch {
	case v.beans < ChosenCoffeeType.beans:
		p("Not enough beans, please fill the container")
		fallthrough
	case v.milk < ChosenCoffeeType.milk:
		p("Not enough beans, please fill the container")
		fallthrough
	case v.water < ChosenCoffeeType.water:
		p("Not enough water, please fill the container")
		fallthrough
	case v.cups <= 0:
		p("Not enough cups, please fill the container")
	default:
		canMakeCupOfCoffee = true
		p("def")
	}

	if canMakeCupOfCoffee {
		pf("Preparing %s, please wait...\n", ChosenCoffeeType.name)
		time.Sleep(3 * time.Second)
		v.beans -= ChosenCoffeeType.beans
		v.milk -= ChosenCoffeeType.milk
		v.water -= ChosenCoffeeType.water
		v.cups -= 1
		v.cash += ChosenCoffeeType.price
		pf("You %s is ready, please take it)))\n", ChosenCoffeeType.name)
	}
}

func (v *VendingMachine) Balance() {
	defer DisplayGreeting()

	fmt.Println("Balance:\n",
		"water: ", v.water, "\n",
		"milk: ", v.milk, "\n",
		"beans:", v.beans, "\n",
		"cups:", v.cups)
}

func (v *VendingMachine) Fill() {
	var water, milk, beans, cups int
	defer DisplayGreeting()

	fmt.Println("write qty of filling water")
	fmt.Fscan(os.Stdin, &water)
	v.water += water
	fmt.Println(v.water)

	fmt.Println("write qty of filling milk")
	fmt.Fscan(os.Stdin, &milk)
	v.milk += milk
	fmt.Println(v.milk)

	fmt.Println("write qty of filling beans")
	fmt.Fscan(os.Stdin, &beans)
	v.beans += beans
	fmt.Println(v.beans)

	fmt.Println("write qty of filling cups")
	fmt.Fscan(os.Stdin, &cups)
	v.cups += cups
	fmt.Println(v.cups)
}

func DisplayGreeting() {
	fmt.Println("Please, choose action: buy, fill, remaning...")
}

func main() {
	defer fmt.Println("Operation not suppoted")

	menu := make(map[int]CupOfCoffee)
	mespresso := CupOfCoffee{
		water: 250,
		milk:  0,
		beans: 16,
		price: 4,
		name:  "espresso",
	}
	mlatte := CupOfCoffee{
		water: 350,
		milk:  75,
		beans: 20,
		price: 7,
		name:  "latte",
	}
	mcapuccino := CupOfCoffee{
		water: 200,
		milk:  100,
		beans: 12,
		price: 6,
		name:  "capuccino",
	}
	menu[1] = mespresso
	menu[2] = mlatte
	menu[3] = mcapuccino

	fmt.Println(menu)

	v := VendingMachine{100, 500, 100, 10, 5}

	scanner := bufio.NewScanner(os.Stdin)
	DisplayGreeting()
	for scanner.Scan() {
		action := scanner.Text()
		switch action {
		case "buy":
			v.Buy(menu)
		case "fill":
			v.Fill()
		case "remaning":
			v.Balance()
		case "cash":
			fmt.Printf("Cash is %d $: \n", v.cash)
			DisplayGreeting()
			//default:
			//  fmt.Println("Operation not suppoted")
		}
	}
}
