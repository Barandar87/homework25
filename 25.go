package main

import (
	"fmt"
	"log"
)

func main() {
	var n string
	fmt.Print("Введите нужные данные: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели такие данные: %v.\n", n)
}
