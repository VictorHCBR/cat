package cat

import (
	grown "github.com/VictorHCBR/adult_cat"
)

func Meow() string {
	return "Meow"
}

func Meows() string {
	return "Meow Meow Meow"
}

func GrownMeow() string {
	return grown.GrownCat(Meow())
}

func GrownMeows() string {
	return grown.GrownCat(Meows())
}