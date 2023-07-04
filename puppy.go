package puppy

import (
	"fmt"

	"github.com/ayushbhatt29/dog"
)

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woof woof"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func from1() {
	fmt.Println("I am from v1.0.0")
}
