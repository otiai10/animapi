package person

type Person struct {
	Name string
}

func New(name string) Person {
	return Person{name}
}
