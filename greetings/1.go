package work

type Person struct {
	Name string
	Age int
}
func (s *Person) GetName() string {
	return s.Name
}

func (s *Person) GetAge() int {
	return s.Age
}