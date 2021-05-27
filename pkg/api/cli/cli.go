package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
	"github.com/gregless22/globalx-sorter-challange/pkg/serializer"
)

type service struct {
	args         []string
	personSerice person.Service
	serializer   serializer.Serializer
}

// NewService creates a new CLI api.  It requires a serialzer to turn []byte  into []Names and the person service / domain
func NewService(args []string, p person.Service, ns serializer.Serializer) error {
	cli := service{args, p, ns}
	if err := cli.checkArgs(); err != nil {
		return err
	}

	if err := cli.runArgs(); err != nil {
		return err
	}

	return nil
}

// check Args will check the validity of the arguments entered
func (s service) checkArgs() error {
	if len(s.args) < 1 {
		return errors.New("no file added as a command line argument")
	}

	if s.args[0] == "" {
		return errors.New("filename cannot be empty")
	}

	return nil
}

// runArgs will parse the command and run the commands and call the functions on the service
func (s service) runArgs() error {
	// only arg at the moment is to read in the file
	// In the future make this a switch statement, parse flags and then call the other functions as required.

	// container for storing the names after scans
	names := []person.Names{}
	filename := "./sorted-names-list.txt"

	// open the file and write to names slice
	if err := s.openFile(&names); err != nil {
		return err
	}

	fmt.Println("starting the sort")
	// call the sorter function form the domain service
	if err := s.personSerice.NameSort(&names); err != nil {
		return err
	}

	if err := s.writeFile(names, filename); err != nil {
		return err
	}

	s.writeStdOut(names)

	return nil

}

// open file will open the file (stored in args on the struct) and unmarshall into []person.Names
func (s service) openFile(names *[]person.Names) error {
	f, err := os.Open(s.args[0])
	if err != nil {
		return errors.New("error opening the file")
	}
	defer f.Close()

	// Read in line by line
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// send to get decoded
		p := &person.Names{}
		s.serializer.Decode(scanner.Bytes(), p)
		*names = append(*names, *p)
	}

	if len(*names) < 1 {
		return errors.New("not able to get names from file")
	}

	return nil
}

func (s service) writeFile(names []person.Names, filename string) error {
	// Open the new file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, n := range names {
		myString := fmt.Sprintf("%v %v\n", n.Firstname, n.Lastname)
		writer.WriteString(myString)

	}
	writer.Flush()

	return nil
}

func (s service) writeStdOut(names []person.Names) {
	for _, n := range names {
		myString := fmt.Sprintf("%v %v\n", n.Firstname, n.Lastname)
		fmt.Print(myString)
	}
}
