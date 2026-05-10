package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/Augajom/road-to-golang/package"
	// "github.com/Augajom/road-to-golang/package/internal/house" // Cannot import internal package in main package
)

type Student struct {
	FirstName  string
	LastName   string
	Age   int
	Grade string
	Address Address
}

// Another struct for demonstration
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Function to demonstrate parameters
func sayHello(name string, round int) {
	for i := 0; i < round; i++{
		fmt.Printf("Hello, %s\n", name)
	}
}

// Method with receiver of type Student
// This method returns the fullname of the student by concatenating the first name and last name.
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

func main() {
	// fullName := "Suphamethee Intharalib" // var fullName = "Suphamethee Intharalib"
	// age := 22 // var age = 22
	// grade := 3.0 // var grade float = 3.0
	id := uuid.New()
	fmt.Println("Hello World")
	fmt.Printf("UUID: %s\n", id)
	// fmt.Printf("Full Name: %s Age: %d Grade: %.1f\n", fullName, age, grade)
	suphamethee.HelloSuphamethee()
	suphamethee.HelloTest()

	// Structs, Arrays
	// var student [3]Student
	// student[0].FirstName = "Suphamethee"
	// student[0].LastName = "Intharalib"
	// student[0].Age = 22
	// student[0].Grade = "A"
	// student[0].Address = Address{
	// 	Street: "123 Main St",
	// 	City:   "Anytown",
	// 	State:  "CA",
	// 	Zip:    "12345",
	// }

	// fmt.Println(student[0])

	// Structs, Maps
	studentMap := make(map[int]Student)
	studentMap[1] = Student{FirstName: "Suphamethee", LastName: "Intharalib", Age: 22, Grade: "A", Address: Address{Street: "123 Main St", City: "Anytown", State: "CA", Zip: "12345"}}
	fmt.Println(studentMap[1])

	// functions parameters
	sayHello("Au", 2)
	sayHello("Go", 5)

	// methods
	student := Student{FirstName: "Suphamethee", LastName: "Intharalib"}

	fullName := student.FullName()
	fmt.Printf("Full Name: %s\n", fullName)
}