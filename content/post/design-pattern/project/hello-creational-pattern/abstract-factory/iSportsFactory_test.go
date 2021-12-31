package abstract_factory

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
