package main

import (
	"errors"
	"fmt"
	"os"
)

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
