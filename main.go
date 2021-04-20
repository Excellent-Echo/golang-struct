package main

import (
	"fmt"
)

type Person struct {
	name   string // " "
	age    int    // 0
	status bool   // false
	// score   []int // slice
	// hobi    map[string]string // map
	// address []Address // struct lain (ditampung di dalam slice)
	// dataTambahan []map[string]Address
}

type Address struct {
	name  string
	jalan string
}

// func Indentity(person Person) string {
// 	age := strconv.Itoa(person.age)

// 	return "hello, " + person.name + " berumur " + age
// }

func createStruct(name string, age int, status bool) Person {
	data := Person{name, age, status}

	return data
}

func main() {
	p1 := createStruct("afista", 23, true)

	fmt.Println(p1)
	// p1 := Person{}

	// p1.score = append(p1.score, 81, 98, 72, 65)

	// p1.hobi = map[string]string{ // inisialisasi
	// 	"ekskul": "pramuka",
	// 	"makan":  "nasi goreng",
	// 	"minum":  "kelapa muda",
	// }

	// p1.hobi["nonton"] = "anime"

	// addr1 := Address{"mahakkam", "waru", "waru-waru", "jawa"}
	// addr2 := Address{"mahakkam", "waru", "waru-waru", "jawa"}

	// p1.address = append(p1.address, addr1, addr2)

	// fmt.Println(p1.address.jalan.name)

	// struct on slice
	// slicePerson := make([]Person, 2, 10)

	// slicePerson[0] = Person{"afis", 23, true}
	// slicePerson[1] = Person{"david", 29, true}

	// slicePerson = append(slicePerson, p1, p2)

	// value := p1.greeting("david")

	// fmt.Println(slicePerson)

	// p1 := Person{"afis", 23, true}
	// p2 := Person{"david", 29, true}

	// var mapStruct = map[string]Person{
	// 	"afista": p1,
	// 	"david":  p2,
	// }

	// for _, value := range mapStruct {
	// 	fmt.Println(value.age)
	// }
}

// type Address struct {
// 	jalan     string
// 	kecamatan string
// 	kabupaten string
// 	provinsi  string
// }

// type Jalan struct {
// 	name string
// 	RT   int
// 	RW   int
// }

// method struct
// func (person Person) greeting(yourName string) string {
// 	return "hello, saya " + person.name + " anda " + yourName
// }

// method yang lain
// func (p1 Person) checkAge() string {
// 	if p1.age < 20 {
// 		return "umur masih muda"
// 	} else {
// 		return "umur sudah dewasa"
// 	}
// }

// struct on map
