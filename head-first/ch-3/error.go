package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	err := fmt.Errorf("a height of %.2f is invalid", -123.4)

	fmt.Println(err.Error())
	fmt.Println(err)

	err = errors.New("error~")
	fmt.Println(err)
	// log.Fatal(err)

	amount, err := paintNeeded(4.2, -3.0)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(amount)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("paintNeeded: a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("paintNeeded: a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
}
