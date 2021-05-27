package person

// Names struct for storing the persons name
type Names struct {
	Firstname string
	Lastname  string
}

type names struct {
	sorter NameSorter
}

// NameSort calls the sorter namesort function
func (s names) NameSort(p *[]Names) error {
	return s.sorter.NameSort(p)
}
