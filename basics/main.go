package main

import "fmt"

func main() {
	fmt.Println("hola mundo")

	// variables
	var whatToSay string
	whatToSay = "adios mundo cruel"
	fmt.Println(whatToSay)

	var i int = 420
	// i = "cat" this is an error
	fmt.Println("i es", i)

	// funciones
	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)

	some, thing := saySomeThings()
	fmt.Println(some, thing)
}

// función básica
func saySomething() string {
	return "something"
}

func saySomeThings() (string, string) {
	return "somethings", "else"
}
