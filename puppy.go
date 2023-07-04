package puppy

import "github.com/ayushbhatt29/dog"

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woof woof"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}
