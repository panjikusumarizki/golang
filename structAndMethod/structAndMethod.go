package structandmethod

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// Struct sebagai properti di struct lain
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// Method
func (user User) display() string {
	return fmt.Sprintf("Nama: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

// Struct sebagai parameter
func displayUser(user User) string {
	return fmt.Sprintf("Nama: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

func displayGroup(group Group) {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name: ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func (group Group) displayGroup() {
	fmt.Printf("Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count: %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name: ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func main() {
	// Cara 3
	user3 := User{3, "Bella", "Attamimi", "bella@mail.com", true}
	result := user3.display()
	fmt.Println(result)

	user4 := User{4, "Hadiza", "Husein", "diza@mail.com", true}
	fmt.Println(user4.display())

	// Pemanggilan struct sebagai properti lain di sebuah struct
	users := []User{user3, user4}

	group := Group{"Gamer", user3, users, true}

	fmt.Println("")

	// displayGroup(group)
	group.displayGroup()

	// Pemanggilan struct sebagai parameter
	// displayUser1 := displayUser(user3)

	// Cara 2
	// user2 := User{
	// 	ID:        2,
	// 	FirstName: "Ervina",
	// 	LastName:  "Mulya",
	// 	Email:     "erv@mail.com",
	// 	IsActive:  true,
	// }

	// Cara 1
	// user1 := User{}
	// user1.ID = 1
	// user1.FirstName = "panji"
	// user1.LastName = "rizki"
	// user1.Email = "mail@mail.com"
	// user1.IsActive = true

	// fmt.Println(displayUser1)
}
