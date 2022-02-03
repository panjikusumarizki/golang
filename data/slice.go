package data

import "fmt"

func main() {
	// Slice
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "PlayStation4")
	gamingConsoles = append(gamingConsoles, "XBox")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	gamingConsoles = append(gamingConsoles, "Android")

	for _, console := range gamingConsoles {
		fmt.Println(console)
	}
}
