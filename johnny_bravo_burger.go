package jhonny_bravo_burger_kata

var ingredientPrices = map[string]int{
	"cheese": 2,
	"bacon":  3,
	"egg":    2,
}

type Burger struct {
	name        string
	price       int
	ingredients []string
}

func CalculateBurgerPrice(burger Burger) Burger {
	price := 5
	for _, element := range burger.ingredients {
		price = price + ingredientPrices[element]
	}

	return Burger{
		name:        burger.name,
		price:       price,
		ingredients: burger.ingredients,
	}
}
