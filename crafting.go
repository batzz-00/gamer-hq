package main

type Craft struct {
}

type Ingredient struct {
	Item   Item
	Amount int
}

type Recipe struct {
	Ingredients []Ingredient
	Result      Item
}
