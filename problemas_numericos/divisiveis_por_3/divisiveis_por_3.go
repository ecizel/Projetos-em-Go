package main

import "fmt"

func main() {

	fmt.Println("Números divisiveis por 3")
	fmt.Println("Entre 1 e 100")

	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Printf("%v é divisível por 3!\n", i)
		}

	}

}
