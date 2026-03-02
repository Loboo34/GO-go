package main

import (
	"fmt"
	"strings"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

var products = []Product{
	{ID: 1, Name: "iPhone 13", Category: "Electronics", Price: 1200},
	{ID: 2, Name: "Samsung TV", Category: "Electronics", Price: 800},
	{ID: 3, Name: "Nike Shoes", Category: "Fashion", Price: 150},
	{ID: 4, Name: "Office Chair", Category: "Furniture", Price: 300},
}

func search(products []Product, query string) []Product {
	var results []Product
	query = strings.ToLower(query)

	for _, p := range products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(p.Category), strings.ToLower(query)) {
			results = append(results, p)
		}
	}

	return results
}

func FindByID(products []Product, id int) *Product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func main() {
	var input string
	fmt.Print("What would you like to search: ")

	fmt.Scan(&input)

	found := search(products, input)

	for _,p := range found{
		fmt.Println(p.Name, p.Category, p.Price)
	}

}
