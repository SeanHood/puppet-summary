//
// Entry-point to the puppet-summary service.
//

package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"os"
)

//
// Setup our sub-commands and use them.
//
func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&serveCmd{}, "")
	subcommands.Register(&pruneCmd{}, "")
	subcommands.Register(&yamlCmd{}, "")
	subcommands.Register(&metricsCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
