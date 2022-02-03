package pointer

import (
	"fmt"
)

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// Pointer struct sebagai parameter
// func graduate(student *Student) {
// 	student.Name = student.Name + " S.Kom"
// }

// Method dengan pointer
func (student *Student) graduate() {
	student.Name = student.Name + " S.Kom"
}

func main() {
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 19

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// Contoh kasus penggunaan pointer
	// number := 5
	// fmt.Println("Alamat memory: ", &number)
	// fmt.Println("Nilai awal: ", number)

	// change(&number, 100)

	// fmt.Println("Alamat memory: ", &number)
	// fmt.Println("Nilai akhir: ", number)

	student := Student{1, "Panji Kusumarizki", 3.40}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory: ", old)
	fmt.Println("Di dalam function: ", *old)
}
