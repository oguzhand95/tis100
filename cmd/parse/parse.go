package parse

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"gopkg.in/yaml.v3"

	"github.com/oguzhand95/tis100/parser"
)

const help = `Parse TIS100 assembly program and output the AST tree`

type Cmd struct {
	Path string `arg:"" name:"path" help:"path to tis100 assembly program"`
}

func (c *Cmd) Run(k *kong.Kong) error {
	b, err := os.ReadFile(c.Path)
	if err != nil {
		return fmt.Errorf("failed to open the file in path %s: %w", c.Path, err)
	}

	ast, err := parser.Parse(bytes.NewReader(b), filepath.Base(c.Path))
	if err != nil {
		return fmt.Errorf("failed to parse the program: %w", err)
	}

	y, err := yaml.Marshal(ast)
	if err != nil {
		return fmt.Errorf("failed to marshal AST of the parsed TIS100 program: %w", err)
	}

	_, err = fmt.Fprint(k.Stdout, string(y))
	if err != nil {
		return fmt.Errorf("failed to print the AST to the stdout: %w", err)
	}

	return nil
}

func (c *Cmd) Help() string {
	return help
}
