package vendingmachine

type config struct {
	menu []Coffee
}

func NewConfig() *config {
	return &config{
		menu: NewMenu(),
	}
}

func NewMenu() []Coffee {
	return []Coffee{
		Coffee{
			id:    1,
			name:  "espresso",
			water: 250,
			milk:  0,
			beans: 16,
			price: 4,
		},
		Coffee{
			id:    2,
			name:  "latte",
			water: 350,
			milk:  75,
			beans: 20,
			price: 7,
		},
		Coffee{
			id:    3,
			name:  "capuccino",
			water: 200,
			milk:  100,
			beans: 12,
			price: 6,
		},
	}
}
