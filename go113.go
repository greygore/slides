package main

import (
	"errors"
	"fmt"
	"os"
)

type wrap interface {
	// START UNWRAP INTERFACE OMIT
	Unwrap() error
	// END UNWRAP INTERFACE OMIT
}

/*
These are just strings we want to put in code blocks:

// START UNWRAP DEFINITION OMIT
errors.Unwrap(err error) error
// END UNWRAP DEFINITION OMIT

// START IS DEFINITION OMIT
errors.Is(err, target error) bool
// END IS DEFINITION OMIT

// START IS EXAMPLE1 OMIT
if errors.Is(err, sql.ErrNoRows) {
	return 0
}
// END IS EXAMPLE1 OMIT

// START IS EXAMPLE2 OMIT
if errors.Is(err, io.EOF) {
	log.Println("Finished reading file")
}
// END IS EXAMPLE2 OMIT


// START AS DEFINITION OMIT
errors.As(err error, target interface{}) bool
// END AS DEFINITION OMIT

// START INDEX PANIC OMIT
runtime error: index out of range [3] with length 1
// END INDEX PANIC OMIT

// START godoc OMIT
go get golang.org/x/tools/cmd/godoc
// END godoc OMIT
*/

func asExample() {
	// START AS EXAMPLE OMIT
	if _, err := os.Open("/does/not/exist"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
	// END AS EXAMPLE OMIT
}

func main() {
	asExample()
}
