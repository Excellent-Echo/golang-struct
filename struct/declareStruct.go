package struct

// membuat struct
type Person struct {
	firstName, lastName, address string // name string dan address string
	age                          int
	status                       bool
}

func DeclareStruct() {
	// fmt.Println(Person{"afista", "pratama", "jember", 23, true})

	// instance
	// p1 := Person{}

	// p1.firstName = "afista"
	// p1.lastName = "pratama"
	// p1.address = "jember"
	// p1.age = 23
	// p1.status = true

	// fmt.Println(p1)

	// langsung diisi (2 cara)
	// 1 langsung diisi sesuai urutan
	// var p1 = Person{"afista", "pratama", "jember", 23, true}

	// fn := "afista"
	// ln := "pratama"

	// 2 diisi sesuai properti
	// var p1 = Person{
	// 	firstName: fn,
	// 	lastName:  ln,
	// 	address:   "jember",
	// 	age:       23,
	// 	status:    true,
	// }
	// fmt.Println(p1)
	fmt.Println("hello world")
}