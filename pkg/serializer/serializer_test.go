package serializer

import (
	"testing"

	"github.com/gregless22/globalx-sorter-challange/pkg/person"
	"github.com/matryer/is"
)

func TestNameSerializer(t *testing.T) {
	testBytes := []byte(`Adonis Julius Archer`)
	p := &person.Names{}
	is := is.New(t)
	testSerializer := NewCLISerializer()
	err := testSerializer.Decode(testBytes, p)
	is.NoErr(err)

	is.Equal(p.Firstname, "Adonis Julius")
	is.Equal(p.Lastname, "Archer")

}
