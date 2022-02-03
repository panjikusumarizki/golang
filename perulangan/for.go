package perulangan

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Cek angka:", i)
	// }

	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Cek angka:", i)
	// 	i++
	// }

	title := "Pergi ke bandung bersama ervina"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, "letter :", string(letter))
		}
	}
}
