package main

import (
	"os"

	"git.sula.io/solevis/poultracker/internal/cmd"
)

func main() {
	// Call Execute so that defers work properly, since os.Exit won't call defers.
	os.Exit(cmd.Execute())
}
