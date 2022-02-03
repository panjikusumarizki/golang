package data

import "fmt"

func main() {
	// Array
	// var languages [5]string

	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Python"
	// languages[3] = "JavaScript"
	// languages[4] = "Java"
	languages := [...]string{
		"Go",
		"Ruby",
		"Python",
		"JavaScript",
		"Java",
		"Kotlin",
	}

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	for index, lang := range languages {
		fmt.Println("index: ", index, ",", "languages: ", lang)
	}
}
