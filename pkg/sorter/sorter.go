package sorter

import (
	"sort"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
)

// Sorter interface takes and array of strings and
// will return them sorted in alphabetical order of the last word

// NewLastNameSorter returns a new last name sorter that addheres to the person.NameSorter interface.
func NewLastNameSorter() person.NameSorter {
	return lastNameSorter{}
}

type lastNameSorter struct{}

// NameSort function will sort the persons by their last name.
func (lastNameSorter) NameSort(n *[]person.Names) error {
	sort.Sort(names(*n))
	return nil
}

// use the golang sort package which requires the Len, Less and Swap interfaces
// This just extends the []person.Names type and allows us to use the golang sort package.
type names []person.Names

func (n names) Len() int {
	return len(n)
}

func (n names) Less(i, j int) bool {
	return n[i].Lastname < n[j].Lastname
}

func (n names) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
