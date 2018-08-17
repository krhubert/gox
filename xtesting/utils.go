package xtesting

import (
	"fmt"
	"os"
)

// ExitOnError checks if err is not nil, print the message and exit.
func ExitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
