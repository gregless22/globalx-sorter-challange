package person

import (
	"testing"

	"github.com/matryer/is"
)

type mockNameSorter struct{}

func (mockNameSorter) NameSort(*[]Names) error {
	return nil
}

func TestNameSort(t *testing.T) {
	is := is.New(t)
	testNameSort := mockNameSorter{}
	testService := NewService(testNameSort)
	p := []Names{}

	err := testService.NameSort(&p)
	is.NoErr(err)

}
