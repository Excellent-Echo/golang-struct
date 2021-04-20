package struct

// embedded struct
type Person struct {
	firstName, lastName string
	age                 int
	status              bool
}

type Student struct {
	Person
	class, batch string
}

func EmbeddedStruct() {
	s1 := Student{}

	s1.firstName = "afista"
	s1.lastName = "pratama"
	s1.class = "skilvul"
	s1.batch = "batch 1"


	fmt.Println(s1.firstName, s1.lastName, s1.class, s1.batch)
}

