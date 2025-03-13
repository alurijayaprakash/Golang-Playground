package main

import "fmt"

func Run() error {
	fmt.Println("Starting our application")
	return nil
}

func main() {
	fmt.Println("Welcome to Comments REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
