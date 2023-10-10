package main

import "fmt"

func print() {
	age := 35
	name := "shaun"
	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line! \n")

	// Println
	fmt.Println("Hello Ninjas")
	fmt.Println("Goodbye Ninjas")
	fmt.Println("My age is ", age, "and my name is", name)

	// Printf ("Formated  strings") %_=formnat  identifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name)
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("Your scored %0.3f points \n", 225.55)

	// Sprintf( save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)

}
