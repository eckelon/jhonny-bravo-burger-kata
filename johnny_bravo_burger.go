package jhonny_bravo_burger_kata

type Ingredient int

const (
	burger Ingredient = iota
	cheese
	bacon
	egg
)

func IngredientPrice(ingredient Ingredient) int {
	switch ingredient {
	case cheese:
		return 2
	case bacon:
		return 3
	case egg:
		return 2
	}
	return 5
}

type Burger struct {
	price       int
	ingredients []Ingredient
}

func hasIngredient(ingredient Ingredient, ingredients []Ingredient) bool {
	hasIt := false
	for _, i := range ingredients {
		hasIt = i == ingredient
	}
	return hasIt
}

func MakeBurger(ingredients []Ingredient) Burger {
	if ingredients == nil || len(ingredients) == 0 {
		return calculateBurgerPrice(Burger{ingredients: []Ingredient{burger}})
	}

	if hasIngredient(burger, ingredients) {
		return calculateBurgerPrice(Burger{ingredients: ingredients})
	}

	return calculateBurgerPrice(Burger{ingredients: append(ingredients, burger)})

}

func calculateBurgerPrice(burger Burger) Burger {
	price := 0
	for _, ingredient := range burger.ingredients {
		price = price + IngredientPrice(ingredient)
	}

	return Burger{
		price:       price,
		ingredients: burger.ingredients,
	}
}
