package main

import (
	"github.com/assanoff/VendingMachineEmulator/internal/vendingmachine"
)

func main() {

	config := vendingmachine.NewConfig()

	vendingmachine.Start(config)
}
