package tests

import (
	"log"
	"tidy/functions"
)

func Test_sum() {
	c := functions.Sum(2, 2)
	if c == 4 {
		log.Print("Sum of 2 + 2 = ", c)
	} else {
		log.Fatal("Error sum of 2 + 2 = ", c)
	}
}
