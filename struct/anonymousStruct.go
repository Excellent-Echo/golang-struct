package struct

type Data struct {
	title       string
	description string
}

func AnonymousStruct() {

	var person = struct {
		name   string
		age    int
		status bool
	}{
		name:   "afis",
		age:    23,
		status: true,
	}

	// person.name = "afis"
	// person.age = 23
	// person.status = true

	person.name = "jaya-jaya"

	fmt.Println(person)
}