package serializer

import (
	"strings"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
)

// Serializer interfaces requires a Decode function
type Serializer interface {
	Decode([]byte, *person.Names) error
}

// CLISerializer returns a new data serializer for the CLI api
func NewCLISerializer() Serializer {
	return cliSerializer{}
}

type cliSerializer struct{}

// Decode func takes and slice of bytes such as those read in from a file and unmarshalls them into the person.Names struct
func (cliSerializer) Decode(b []byte, p *person.Names) error {
	myString := string(b)

	// split to get firstname and lastname
	strArray := strings.Fields(myString)
	p.Firstname = strings.Join(strArray[:len(strArray)-1], " ")
	p.Lastname = strArray[len(strArray)-1]

	return nil
}
