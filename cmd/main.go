package main

import (
	"github.com/alecthomas/kong"

	"github.com/oguzhand95/tis100/cmd/root"
)

func main() {
	cli := &root.Cli{}
	ctx := kong.Parse(cli,
		kong.Name("tis100"),
		kong.Description("A command line interface to do TIS100 related jobs"),
		kong.UsageOnError(),
	)

	ctx.FatalIfErrorf(ctx.Run())
}
