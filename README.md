# TIS100

A tool and package to parse TIS100 assembly programs.

## CLI Usage

### Command `tis100 parse <path>`

Parses the TIS100 assembly program in the given `<path>`, and prints the AST result into standard output.

```
MOV 10, ACC
MOV ACC, DOWN
```

```bash
-> tis100 parse ~/program.asm
```

The result for the above execution looks like this;

```yaml
instructions:
  - mov:
      source:
        literal: 10
      destination:
        register: ACC
  - mov:
      source:
        register: ACC
      destination:
        register: DOWN
```

## Package Usage

```bash
go get github.com/oguzhand95/tis100
```

```go
import "github.com/oguzhand95/tis100/parser"

func main() {
    b, err := os.ReadFile(c.Path)
    if err != nil {
        return fmt.Errorf("failed to open the file in path %s: %w", c.Path, err)
    }

    ast, err := parser.Parse(bytes.NewReader(b), filepath.Base(c.Path))
    if err != nil {
        return fmt.Errorf("failed to parse the program: %w", err)
    }
	
    // Use ast struct to some stuff 
}
```

# Thanks

- Thanks [Zachtronics](https://www.zachtronics.com/) for developing this awesome [puzzle game](https://store.steampowered.com/app/370360/TIS100/)! 