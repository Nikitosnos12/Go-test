package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("другой способ")

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

}
