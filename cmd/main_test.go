package main

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

// TestCommandLineArg test that if a command line argument is passed in
func TestCommandLineArg(t *testing.T) {
	is := is.New(t)

	args := []string{"./testFile"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	// fmt.Println(err.Error())

	// out := stdout.String()
	// is.True(strings.Contains(out, "./testFile"))

}

// TestNoArg will check if no file is put in we get an error back
func TestNoArg(t *testing.T) {
	is := is.New(t)

	args := []string{}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.True(err != nil)
}
