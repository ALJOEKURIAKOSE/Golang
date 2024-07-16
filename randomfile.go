package main

import (
	"fmt"
)

func main() {
	// var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	cars := map[string]map[string]interface{}{}

	cars["Volvo"] = map[string]interface{}{
		"price":   1,
		"mileage": 100}

	cars["Volvo"]["price"] = 1200
	fmt.Print(cars["Volvo"]["price"])
	// fmt.Printf("random file ")





}
