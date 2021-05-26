package person

// Service interface defines what packages the domain requires to work. It needs a Namesorter
type Service interface {
	NameSort(*[]Names) error
	// Encode([]Person) ([]byte, error)
	// Decode([]byte) (chan Person, error)
}

// NameSorter interface requires the function Namesort which will sort the names by their last name
type NameSorter interface {
	NameSort(*[]Names) error
}

// Serializer is to Encode and Decode though can do this by Marshalling
type Serializer interface {
	Encode([]Names) ([]byte, error)
	Decode([]byte) (chan Names, error)
}
