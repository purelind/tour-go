package main

import (
	"log"
	"os"

	_ "github.com/purelind/tour-go/goinaction/chapter2/sample/matchers"
	"github.com/purelind/tour-go/goinaction/chapter2/sample/search"
)

// init is called prior to main
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}