package fungsi

import "fmt"

func main() {
	result := add(10, 30)
	fmt.Println(result)

	luas, keliling := calculate(9, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func add(number, numberTwo int) int {
	return number + numberTwo
}

// function dengan multiple return
func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// function dengan predefined return value
func calculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return luas, keliling
}
