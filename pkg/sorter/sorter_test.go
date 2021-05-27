package sorter

import (
	"testing"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
	"github.com/matryer/is"
)

// TestSortFunction tests the entire name sorter.
func TestSortFunction(t *testing.T) {
	is := is.New(t)

	testSorter := NewLastNameSorter()
	testArray := []person.Names{
		{Firstname: "Janet", Lastname: "Parsons"},
		{Firstname: "Vaughn", Lastname: "Lewis"},
		{Firstname: "Adonis Julius", Lastname: "Archer"},
		{Firstname: "Shelby Nathan", Lastname: "Yoder"},
		{Firstname: "Marin", Lastname: "Alvarez"},
		{Firstname: "London", Lastname: "Lindsey"},
		{Firstname: "Beau Tristan", Lastname: "Bentley"},
		{Firstname: "Leo", Lastname: "Gardner"},
		{Firstname: "Hunter Uriah Mathew", Lastname: "Clarke"},
		{Firstname: "Mikayla", Lastname: "Lopez"},
		{Firstname: "Frankie Conner", Lastname: "Ritter"},
	}

	resultArray := []person.Names{
		{Firstname: "Marin", Lastname: "Alvarez"},
		{Firstname: "Adonis Julius", Lastname: "Archer"},
		{Firstname: "Beau Tristan", Lastname: "Bentley"},
		{Firstname: "Hunter Uriah Mathew", Lastname: "Clarke"},
		{Firstname: "Leo", Lastname: "Gardner"},
		{Firstname: "Vaughn", Lastname: "Lewis"},
		{Firstname: "London", Lastname: "Lindsey"},
		{Firstname: "Mikayla", Lastname: "Lopez"},
		{Firstname: "Janet", Lastname: "Parsons"},
		{Firstname: "Frankie Conner", Lastname: "Ritter"},
		{Firstname: "Shelby Nathan", Lastname: "Yoder"},
	}

	// //  Test no errors
	err := testSorter.NameSort(&testArray)
	is.NoErr(err)

	// // Test array is the same (may need to write better comparison test)
	is.Equal(testArray, resultArray)

}

// TestLen will test the length
func TestLen(t *testing.T) {
	is := is.New(t)
	testArray := names{
		{Firstname: "Janet", Lastname: "Parsons"},
		{Firstname: "Vaughn", Lastname: "Lewis"},
		{Firstname: "Adonis Julius", Lastname: "Archer"},
	}

	length := testArray.Len()
	is.Equal(length, 3)
}

// TestLess will test that we are comparing things in the correct order.  We expect the Parsons > Lewis as it is later in the alphabet
func TestLess(t *testing.T) {
	is := is.New(t)
	testArray := names{
		{Firstname: "Janet", Lastname: "Parsons"},
		{Firstname: "Vaughn", Lastname: "Lewis"},
	}

	order := testArray.Less(0, 1)
	is.Equal(order, false)
}

// TestSwap will make sure the positions are swapped correctly
func TestSwap(t *testing.T) {
	is := is.New(t)
	testArray := names{
		{Firstname: "Janet", Lastname: "Parsons"},
		{Firstname: "Vaughn", Lastname: "Lewis"},
		{Firstname: "Adonis Julius", Lastname: "Archer"},
	}

	swapArray := names{
		{Firstname: "Adonis Julius", Lastname: "Archer"},
		{Firstname: "Vaughn", Lastname: "Lewis"},
		{Firstname: "Janet", Lastname: "Parsons"},
	}

	// Swap positions of two objects, could randomise this for a better test
	testArray.Swap(0, 2)
	is.Equal(testArray, swapArray)
}
