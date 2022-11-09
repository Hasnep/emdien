package main

import (
	"os"

	"github.com/Hasnep/emdien/emdien"
)

func main() {
	err := emdien.Execute()
	if err != nil {
		os.Exit(1)
	}
}
