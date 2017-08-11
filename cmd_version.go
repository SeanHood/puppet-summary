//
// Show our version - This uses a level of indirection for our test-case
//

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io"
	"os"
)

//
// modified during testing
//
var out io.Writer = os.Stdout
var exit func(code int) = os.Exit

var (
	version string = "unreleased"
)

type versionCmd struct{}

//
// Glue
//
func (*versionCmd) Name() string     { return "version" }
func (*versionCmd) Synopsis() string { return "Show our version." }
func (*versionCmd) Usage() string {
	return `version :
  Report upon our version, and exit.
`
}

//
// Flag setup
//
func (p *versionCmd) SetFlags(f *flag.FlagSet) {
}

//
// Show the version - using the "out"-writer.
//
func showVersion() {
	fmt.Fprintf(out, "%s\n", version)
}

//
// Entry-point.
//
func (p *versionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	showVersion()
	return subcommands.ExitSuccess
}
