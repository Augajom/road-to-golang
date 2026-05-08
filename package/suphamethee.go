package suphamethee

import (
	"fmt"
	"github.com/Augajom/road-to-golang/package/internal/house"
)

func HelloSuphamethee() {
	fmt.Println("Hello Suphamethee")
	// Can call function within the same package Even though they are different files
	generateTest()
	HelloTest()

	// Can call function from internal package
	house.HelloAddress()
}