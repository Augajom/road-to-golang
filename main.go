package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/Augajom/road-to-golang/package"
	// "github.com/Augajom/road-to-golang/package/internal/house" // Cannot import internal package in main package
)

func main() {
	id := uuid.New()
	fmt.Println("Hello World")
	fmt.Printf("UUID: %s\n", id)
	suphamethee.HelloSuphamethee()
	suphamethee.HelloTest()
}