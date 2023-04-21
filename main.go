package main

import (
	"fmt"
	"os"
)

// Product represents a simple product structure
type Product struct {
	ID    string
	Name  string
	Price float64
}

// CartItem represents an item in the shopping cart
type CartItem struct {
	Product  Product
	Quantity int
}

// In-memory storage for products
var products = []Product{
	{ID: "1", Name: "Laptop", Price: 999.99},
	{ID: "2", Name: "Smartphone", Price: 499.99},
	{ID: "3", Name: "Headphones", Price: 99.99},
}

// Shopping cart
var cart []CartItem

func main() {
	for {
		fmt.Println("\n=== GoShop CLI ===")
		fmt.Println("1. View all products")
		fmt.Println("2. View product details")
		fmt.Println("3. Add to cart")
		fmt.Println("4. View cart")
		fmt.Println("5. Checkout")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option (1-6): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			viewAllProducts()
		case 2:
			viewProductDetails()
		case 3:
			addToCart()
		case 4:
			viewCart()
		case 5:
			checkout()
		case 6:
			fmt.Println("Thank you for using GoShop!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func viewAllProducts() {
	fmt.Println("\n=== All Products ===")
	for _, product := range products {
		fmt.Printf("ID: %s\n", product.ID)
		fmt.Printf("Name: %s\n", product.Name)
		fmt.Printf("Price: $%.2f\n", product.Price)
		fmt.Println("-------------------")
	}
}

func viewProductDetails() {
	fmt.Print("\nEnter product ID: ")
	var id string
	fmt.Scanln(&id)

	for _, product := range products {
		if product.ID == id {
			fmt.Println("\n=== Product Details ===")
			fmt.Printf("ID: %s\n", product.ID)
			fmt.Printf("Name: %s\n", product.Name)
			fmt.Printf("Price: $%.2f\n", product.Price)
			return
		}
	}
	fmt.Println("Product not found!")
}

func addToCart() {
	fmt.Print("\nEnter product ID to add to cart: ")
	var id string
	fmt.Scanln(&id)

	var selectedProduct Product
	found := false
	for _, product := range products {
		if product.ID == id {
			selectedProduct = product
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Product not found!")
		return
	}

	fmt.Print("Enter quantity: ")
	var quantity int
	fmt.Scanln(&quantity)

	if quantity <= 0 {
		fmt.Println("Quantity must be greater than 0")
		return
	}

	// Check if product already in cart
	for i, item := range cart {
		if item.Product.ID == id {
			cart[i].Quantity += quantity
			fmt.Printf("Updated quantity of %s in cart\n", selectedProduct.Name)
			return
		}
	}

	// Add new item to cart
	cart = append(cart, CartItem{
		Product:  selectedProduct,
		Quantity: quantity,
	})
	fmt.Printf("Added %d %s(s) to cart\n", quantity, selectedProduct.Name)
}

func viewCart() {
	if len(cart) == 0 {
		fmt.Println("\nYour cart is empty!")
		return
	}

	fmt.Println("\n=== Shopping Cart ===")
	total := 0.0
	for _, item := range cart {
		itemTotal := item.Product.Price * float64(item.Quantity)
		fmt.Printf("Product: %s\n", item.Product.Name)
		fmt.Printf("Quantity: %d\n", item.Quantity)
		fmt.Printf("Price per item: $%.2f\n", item.Product.Price)
		fmt.Printf("Total for this item: $%.2f\n", itemTotal)
		fmt.Println("-------------------")
		total += itemTotal
	}
	fmt.Printf("Cart Total: $%.2f\n", total)
}

func checkout() {
	if len(cart) == 0 {
		fmt.Println("\nYour cart is empty!")
		return
	}

	viewCart()
	fmt.Print("\nProceed to checkout? (y/n): ")
	var confirm string
	fmt.Scanln(&confirm)

	if confirm == "y" || confirm == "Y" {
		total := 0.0
		for _, item := range cart {
			total += item.Product.Price * float64(item.Quantity)
		}
		fmt.Printf("\nThank you for your purchase!\nTotal amount: $%.2f\n", total)
		cart = nil // Clear the cart
	} else {
		fmt.Println("Checkout cancelled")
	}
}
