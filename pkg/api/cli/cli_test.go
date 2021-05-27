package cli

import (
	"os"
	"testing"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
	"github.com/matryer/is"
)

// mock the personservice
type mockPersonSevice struct{}

func (mockPersonSevice) NameSort(*[]person.Names) error {
	return nil
}

type mockSerializer struct{}

func (mockSerializer) Decode([]byte, *person.Names) error {
	return nil
}

func TestNoArg(t *testing.T) {
	is := is.New(t)
	args := []string{}
	personSevice := mockPersonSevice{}
	serializer := mockSerializer{}

	testService := service{args, personSevice, serializer}

	err := testService.checkArgs()
	is.Equal(err.Error(), "no file added as a command line argument")
}

func TestArgEmpty(t *testing.T) {
	is := is.New(t)
	args := []string{""}
	personSevice := mockPersonSevice{}
	serializer := mockSerializer{}

	testService := service{args, personSevice, serializer}
	err := testService.checkArgs()
	is.Equal(err.Error(), "filename cannot be empty")
}

func TestOpenFile(t *testing.T) {
	is := is.New(t)
	args := []string{"unfound file"}
	personSevice := mockPersonSevice{}
	serializer := mockSerializer{}
	names := &[]person.Names{}

	testService := service{args, personSevice, serializer}
	err := testService.openFile(names)
	is.Equal(err.Error(), "error opening the file")

	// try with correct file
	testService.args = []string{"./test-empty-names-list.txt"}
	err = testService.openFile(names)
	is.Equal(err.Error(), "not able to get names from file")

	// try with correct file
	testService.args = []string{"./test-names-list.txt"}
	err = testService.openFile(names)
	is.NoErr(err)
}

func TestWriteFile(t *testing.T) {
	is := is.New(t)
	args := []string{""}
	personSevice := mockPersonSevice{}
	serializer := mockSerializer{}
	names := []person.Names{{Firstname: "Janet", Lastname: "Parsons"},
		{Firstname: "Vaughn", Lastname: "Lewis"},
		{Firstname: "Adonis Julius", Lastname: "Archer"}}

	testService := service{args, personSevice, serializer}
	err := testService.writeFile(names, ".test-output.txt")
	is.NoErr(err)
}

func TestRunArgs(t *testing.T) {
	is := is.New(t)
	args := []string{"./test-names-list.txt"}
	personSevice := mockPersonSevice{}
	serializer := mockSerializer{}

	testService := service{args, personSevice, serializer}
	err := testService.runArgs()
	is.NoErr(err)
}

func TestCleanUP(t *testing.T) {
	os.Remove("./sorted-names-list.txt")
}
