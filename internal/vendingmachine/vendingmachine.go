package vendingmachine

import (
	"fmt"
	"os"
	"time"
)

// Coffee ...
type Coffee struct {
	id    int
	name  string
	water int
	milk  int
	beans int
	price int
}

type vendingmachine struct {
	water   int
	milk    int
	beans   int
	cash    int
	cups    int
	coins   int
	item    int
	actions map[string]stateFn
	menu    []Coffee
}

type stateFn func(*vendingmachine) stateFn

// Start ...
func Start(config *config) error {

	fmt.Println("starting new vending machine...")

	actions := make(map[string]stateFn)
	actions["exit"] = exitFn
	actions["fill"] = fillFn
	actions["buy"] = buyFn
	actions["remaning"] = balanceFn

	v := &vendingmachine{
		1000,
		5000,
		1000,
		100,
		5,
		0,
		0,
		actions,
		config.menu}

	v.run()
	return nil
}

func (v *vendingmachine) run() {
	for state := displayGreetingFn; state != nil; {
		state = state(v)
	}
}

func displayGreetingFn(v *vendingmachine) stateFn {
	var userInput string
	fmt.Println("Please, select action: buy, fill, remaning, exit...")

	fmt.Fscan(os.Stdin, &userInput)

	stateFn, ok := v.actions[userInput]
	if !ok {
		fmt.Println("Unsupported operation, try again later")
		return displayGreetingFn
	}

	return stateFn
}

func balanceFn(v *vendingmachine) stateFn {

	fmt.Println("Balance:\n",
		"water: ", v.water, "\n",
		"milk: ", v.milk, "\n",
		"beans: ", v.beans, "\n",
		"cups: ", v.cups, "\n",
		"cash:", v.cash)

	return displayGreetingFn
}

func exitFn(v *vendingmachine) stateFn {
	fmt.Println("Goodbye")
	return nil
}

func resetFn(v *vendingmachine) stateFn {
	fmt.Println("Please take your turn")

	v.coins = 0
	v.item = 0
	return displayGreetingFn
}

func fillFn(v *vendingmachine) stateFn {
	var water, milk, beans, cups int

	fmt.Println("inserte qty of filling water")
	fmt.Fscan(os.Stdin, &water)
	v.water += water
	fmt.Println(v.water)

	fmt.Println("inserte qty of filling milk")
	fmt.Fscan(os.Stdin, &milk)
	v.milk += milk
	fmt.Println(v.milk)

	fmt.Println("inserte qty of filling beans")
	fmt.Fscan(os.Stdin, &beans)
	v.beans += beans
	fmt.Println(v.beans)

	fmt.Println("insert qty of filling cups")
	fmt.Fscan(os.Stdin, &cups)
	v.cups += cups
	fmt.Println(v.cups)

	return displayGreetingFn
}

func buyFn(v *vendingmachine) stateFn {
	var id int

	fmt.Println("Please insert number of select item...")

	for _, item := range v.menu {
		fmt.Printf("%d : %s\n", item.id, item.name)
	}
	fmt.Fscan(os.Stdin, &id)
	v.item = id
	selectedItem := v.menu[id-1]
	fmt.Printf("You selected %s, price %d\n", selectedItem.name, selectedItem.price)

	if id <= 0 || id > v.menu[len(v.menu)-1].id {
		fmt.Println("You selected invalid number of item. Try again later...")
		return displayGreetingFn
	}

	return insertMoney
}

func insertMoney(v *vendingmachine) stateFn {

	fmt.Println("Please insert money. You have 10 sec...")

	time.AfterFunc(10*time.Second, func() {
		if v.item != 0 {
			fmt.Println("Times is up. You inserted not enough monye. Try again later")
			resetFn(v)
		}
	})

	sum, total := 0, 0
	selectedItem := v.menu[v.item-1]
	price := selectedItem.price

	for total < price {
		fmt.Fscan(os.Stdin, &sum)
		total += sum
		fmt.Printf("You inserted %d, please add more %d\n", total, price-total)
	}

	if total > price {
		fmt.Printf("Please take your turn %d\n", total-price)
	}

	return processCoffee
}
func processCoffee(v *vendingmachine) stateFn {
	selectedItem := v.menu[v.item-1]
	fmt.Println("Processing coffee:", selectedItem.name)
	v.item = 0

	switch {
	case selectedItem.water > v.water:
		fmt.Println("Not enough water, please fill the container")
		return displayGreetingFn
	case selectedItem.milk > v.milk:
		fmt.Println("Not enough milk, please fill the container")
		return displayGreetingFn
	case selectedItem.beans > v.beans:
		fmt.Println("Not enough beans, please fill the container")
		return displayGreetingFn
	case v.cups < 1:
		fmt.Println("Not enough cups, please fill the container")
		return displayGreetingFn
	case selectedItem.beans > v.beans:
		fmt.Println("Not enough beans, please fill the container")
		return displayGreetingFn
	default:
		v.water -= selectedItem.water
		v.milk -= selectedItem.milk
		v.beans -= selectedItem.beans
		v.cash += selectedItem.price
		v.cups -= 1
		fmt.Printf("Please take you drink %s and good day\n", selectedItem.name)
	}

	return displayGreetingFn
}
