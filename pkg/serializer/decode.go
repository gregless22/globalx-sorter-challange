package serializer

import (
	"strings"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
)

type Serializer interface {
	// Encode([]byte, person.Person)
	Decode([]byte, *person.Names)
}

func NewCLISerializser() Serializer {
	return cliSerialiser{}
}

type cliSerialiser struct{}

func (cliSerialiser) Decode(b []byte, p *person.Names) {
	myString := string(b)

	// split to get firstname and lastname
	strArray := strings.Fields(myString)
	p.Firstname = strings.Join(strArray[:len(strArray)-1], " ")
	p.Lastname = strArray[len(strArray)-1]

}
