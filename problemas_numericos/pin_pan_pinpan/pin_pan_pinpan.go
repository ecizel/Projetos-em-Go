package main

import "fmt"

func main() {

	fmt.Println("Brincadeira Pin, Pan, Pinpan")
	fmt.Println("Sempre que o número for múltiplo de 3: Pin!")
	fmt.Println("Se o número for múltiplo de 5: Pan!")
	fmt.Println("E se for múltiplo dos dois: Pinpan!")
	fmt.Println("Entre 1 e 100")

	for i := 1; i <= 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("Pinpan!")
		case i%3 == 0:
			fmt.Println("Pin!")
		case i%5 == 0:
			fmt.Println("Pan!")
		default:
			fmt.Println(i)
		}

	}

}
