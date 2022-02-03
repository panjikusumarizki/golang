package data

import "fmt"

func main() {
	// Map
	// var myMap = map[string]int{}

	// myMap["JavaScript"] = 9
	// myMap["Go"] = 8
	// myMap["Kotlin"] = 8

	myMap := map[string]string{
		"JavaScript": "is beautiful",
		"Go":         "is super fast",
		"Java":       "is confusing",
	}

	// for key, value := range myMap {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	// fmt.Println("==========")

	// delete(myMap, "Java")

	// for key, value := range myMap {
	// 	fmt.Println("key: ", key, "value: ", value)
	// }

	value, isAvailable := myMap["kotlin"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}
