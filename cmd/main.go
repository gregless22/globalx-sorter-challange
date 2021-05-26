package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gregless22/globalx-sorter-challange/pkg/api/cli"
	"github.com/gregless22/globalx-sorter-challange/pkg/person"
	"github.com/gregless22/globalx-sorter-challange/pkg/serializer"
	"github.com/gregless22/globalx-sorter-challange/pkg/sorter"
)

const (
	// exitFail is the exit code if the program
	// fails.
	exitFail = 1
)

func main() {

	log.Println("NameSorter Starting")
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
	log.Println("NameSorter Ended")
}

// the run function abtracts away from the main func so we can run tests on it
func run(args []string, stdout io.Writer) error {

	nameSorter := sorter.NewLastNameSorter()

	// fileHandler := fileHandler.NewService()
	ps := person.NewService(nameSorter)
	ns := serializer.NewCLISerializser()
	// se := serializser.NewService()
	err := cli.NewService(args[1:], ps, ns)
	if err != nil {
		return err
	}

	// Read in the file

	// decode file to array of persons

	// sort by last name

	// encode to file output and stdout

	return nil
}
