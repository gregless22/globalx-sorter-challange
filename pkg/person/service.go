package person

// Service interface defines what packages the domain requires to work. It needs a Namesorter
type Service interface {
	NameSort(*[]Names) error
}

// NameSorter interface requires the function Namesort which will sort the names by their last name
type NameSorter interface {
	NameSort(*[]Names) error
}

//  Returns a new service for the Person Domain
func NewService(sorter NameSorter) Service {
	return names{sorter}
}
