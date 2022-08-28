package parser

import (
	"fmt"
	"io"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

const lookahead = 2

var l = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Ident", Pattern: `[A-Z]+`},
	{Name: "Colon", Pattern: `:`},
	{Name: "Literal", Pattern: `([-]?[0-9]+)`},
	{Name: "ParamSep", Pattern: `,([ ]+)?`},
	{Name: "Whitespace", Pattern: `[ \t]+`},
	{Name: "EOL", Pattern: `[\n\r]+`},
})

func Parse(r io.Reader, name string) (*Program, error) {
	p, err := participle.Build[Program](
		participle.Lexer(l),
		participle.UseLookahead(lookahead),
		participle.Elide("EOL", "Whitespace"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a parser: %w", err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read all bytes from the given reader: %w", err)
	}

	ast, err := p.ParseBytes(name, b)
	if err != nil {
		return nil, fmt.Errorf("failed to parse from bytes: %w", err)
	}

	return ast, nil
}
