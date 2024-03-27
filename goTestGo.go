package main

import (
	"fmt"

	"github.com/lordofthemind/gormGinGo/helpers"
)

func main() {
	randGen := helpers.NewRandomGenerator()

	fmt.Println("Random Int: ", randGen.RandomInt(1, 100))
	fmt.Println("Random String: ", randGen.RandomString(10))
	fmt.Println("Random Username: ", randGen.RandomUsername())
	fmt.Println("Random Email: ", randGen.RandomEmail())
	fmt.Println("Random Phone: ", randGen.RandomPhone())
	fmt.Println("Random Name: ", randGen.RandomName())
	fmt.Println("Random Address: ", randGen.RandomAddress())
	fmt.Println("Random Age: ", randGen.RandomAge())

}
