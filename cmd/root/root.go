package root

import "github.com/oguzhand95/tis100/cmd/parse"

type Cli struct {
	Parse parse.Cmd `cmd:"" name:"parse"`
}
