package puppy

import "github.com/niutingyuan/cat"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return cat.WhenGrownUp(Bark())
}
func BigBarks() string {
	return cat.WhenGrownUp(Barks())
}
