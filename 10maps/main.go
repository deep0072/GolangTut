package main

import "fmt"

func main() {
	fmt.Println("welcome to maps in golang")

	// An Empty map ==> map[Key_Type]Value_Type{}

	languages := make(map[string]string)

	languages["Py"] = "python"
	languages["js"] = "javascript"
	languages["RB"] = "ruby"

	fmt.Println(languages["RB"])

	//delete the keys in map use delete(map, key_name)

	delete(languages, "Py")

	fmt.Println(languages)

	// loop through the maps
	for key, value := range languages {
		fmt.Println(key, value)
	}
}
