package main

import (
	"fmt"

	"math/rand"

	"github.com/atotto/clipboard"
)

// starcount returns the number of stars to generate
func starcount() int {
	var starcount int
	fmt.Println("How many stars?")
	fmt.Scanln(&starcount)
	return starcount
}

// gen returns a random position and size for a star
func gen() (int, int, float32) {
	// -90 -60 -> 90 60
	posx := rand.Intn(180) - 90
	posy := rand.Intn(120) - 60
	size := float32(rand.Intn(3))
	if size == 0 {
		size = 1
	}
	return posx, posy, size - 0.5
}

// genStars generates the star string
func genStars() string {
	starcount := starcount()
	output := ""

	for i := 0; i < starcount-1; i++ {
		posx, posy, size := gen()
		output += fmt.Sprintf("  star %d %d %f,\n", posx, posy, size)
	}
	posx, posy, size := gen()
	output += fmt.Sprintf("  star %d %d %f", posx, posy, size)

	return output
}

// copyToClip copies the string to the clipboard
func copyToClip(stringToCopy string) {
	err := clipboard.WriteAll(stringToCopy)

	if err != nil {
		panic(err)
	}

	fmt.Println("Copied to clipboard")
}

func main() {
	boilerplate := `
star x y z = group [
  circle z
    |> filled white
    |> move (x, y)]

stars = group [` + genStars() + `] 

myShapes model = [
    rect 200 200
        |> filled (rgb 0 0 0),
    stars]`

	copyToClip(boilerplate)
}
