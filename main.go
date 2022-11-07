package main

import (
	"os"

	"github.com/Hasnep/test-repo-please-delete/emdien"
)

func main() {
	err := emdien.Execute()
	if err != nil {
		os.Exit(1)
	}
}
