package person

type Names struct {
	Firstname string
	Lastname  string
}

//  Leave this here for now but will probably write it in a decoder / encoder packer
// func (p Person) UnmarshalText(text []byte) error {
// 	return nil
// }

// func (p Person) MarshalText() ([]byte, error) {
// 	return []byte(""), nil
// }

func NewService(sorter NameSorter) Service {
	return service{sorter}
}

type service struct {
	sorter NameSorter
	// serializer Serializer
}

// func (s service) Encode(p []Person) ([]byte, error) {
// 	return (s.serializer.Encode(p))
// }

// func (s service) Decode(b []byte) (chan Person, error) {
// 	return (s.serializer.Decode(b))
// }

func (s service) NameSort(p *[]Names) error {
	return s.sorter.NameSort(p)
}
