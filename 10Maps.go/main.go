package main

import "fmt"

func main() {
	fmt.Println("Understanding maps in GO")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	//deleting a pair from the map - we use delete keyword for that and this can be used for deleting an element from the slices as well
	delete(languages, "PY")
	fmt.Println("Modified list of languages: ", languages)

	//looping over a map in GO
	for key, value := range languages {
		fmt.Printf("For key %v, the value is %v\n", key, value)
	}

}
