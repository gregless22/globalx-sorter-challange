package person

type Names struct {
	Firstname string
	Lastname  string
}

type names struct {
	sorter NameSorter
}

func (s names) NameSort(p *[]Names) error {
	return s.sorter.NameSort(p)
}
