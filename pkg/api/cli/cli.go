package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
	"github.com/gregless22/globalx-sorter-challange/pkg/serializer"
)

// type Service interface {
// 	FileReader() ([]byte, error)
// 	// FileWriter() ([]byte error)
// }

type service struct {
	args         []string
	personSerice person.Service
	serializer   serializer.Serializer
}

type FileHandler interface {
	Read(string, *[]person.Names) error
	Write(string, []person.Names) error
}

func NewService(args []string, p person.Service, ns serializer.Serializer) error {
	cli := service{args, p, ns}
	if err := cli.checkArgs(args); err != nil {
		return err
	}

	if err := cli.runArgs(args); err != nil {
		return err
	}

	return nil
}

// check Args will check the validity of the arguments entered
func (service) checkArgs(args []string) error {
	if len(args) < 1 {
		return errors.New("no file added as a command line argument")
	}

	if args[0] == "" {
		return errors.New("filename cannot be empty")
	}

	return nil
}

// runArgs will parse the command and run the commands and call the functions on the service
func (s service) runArgs(args []string) error {
	// only arg at the moment is to read in the file
	// In the future make this a switch statement, parse flags and then call the other functions as required.
	// convert the file to []bytes
	f, err := os.Open(args[0])
	if err != nil {
		return errors.New("error opening the file")
	}
	defer f.Close()

	// Read in line by line
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// container for storing the names after scans
	names := []person.Names{}

	for scanner.Scan() {
		// send to get decoded
		p := &person.Names{}
		s.serializer.Decode(scanner.Bytes(), p)
		names = append(names, *p)
	}
	fmt.Println("starting the sort")
	// call the sorter function form the domain service
	if err := s.personSerice.NameSort(&names); err != nil {
		return err
	}

	// Open the new file
	file, err := os.Create("./sorted-names-list.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, n := range names {
		myString := fmt.Sprintf("%v %v\n", n.Firstname, n.Lastname)
		writer.WriteString(myString)
		fmt.Print(myString)
	}
	writer.Flush()

	return nil

}
