package vendingmachine

import (
	"fmt"
	"os"
)

type Vendingmachine struct {
	water   int
	milk    int
	beans   int
	cash    int
	cups    int
	input   string
	actions map[string]stateFn
}

// type cuoOfCoffee struct {
// 	water, milk, beans, cash, cups int
// 	name                           string
// 	idx                            int
// }

type stateFn func(*Vendingmachine) stateFn

// Start ...
func Start() {

	actions := make(map[string]stateFn)
	actions["exit"] = exitFn

	userInput := "..."

	v := &Vendingmachine{100, 500, 100, 10, 5, userInput, actions}
	v.run()
	return
}

func (v *Vendingmachine) run() {

	for state := displayGreetingFn; state != nil; {
		state = state(v)
	}
}

func displayGreetingFn(v *Vendingmachine) stateFn {
	fmt.Fscan(os.Stdin, &v.input)

	stateFn, ok := v.actions[v.input]
	if !ok {
		return displayGreetingFn
	}

	return stateFn
}

func exitFn(v *Vendingmachine) stateFn {
	fmt.Println("goodbye")
	return nil
}
