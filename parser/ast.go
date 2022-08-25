package parser

type Program struct {
	Instructions []*Instruction `yaml:"instructions" parser:"@@*"`
}

type Instruction struct {
	SimpleOp *SimpleOp `yaml:"simpleOp,omitempty" parser:"( @('NOP'|'SWP'|'SAV'|'NEG')"`
	Add      *Add      `yaml:"add,omitempty" parser:"| 'ADD' @@"`
	Mov      *Mov      `yaml:"mov,omitempty" parser:"| 'MOV' @@"`
	Jmp      *Jmp      `yaml:"jmp,omitempty" parser:"| 'JMP' @@"`
	Jez      *Jez      `yaml:"jez,omitempty" parser:"| 'JEZ' @@"`
	Jnz      *Jnz      `yaml:"jnz,omitempty" parser:"| 'JNZ' @@"`
	Jgz      *Jgz      `yaml:"jgz,omitempty" parser:"| 'JGZ' @@"`
	Jlz      *Jlz      `yaml:"jlz,omitempty" parser:"| 'JLZ' @@"`
	Jro      *Jro      `yaml:"jro,omitempty" parser:"| 'JRO' @@"`
	Sub      *Sub      `yaml:"sub,omitempty" parser:"| 'SUB' @@"`
	Label    string    `yaml:"label,omitempty" parser:"| @(Ident) Colon)"`
}

type Mov struct {
	Source      *Param `yaml:"source" parser:"@@ ParamSep"`
	Destination *Param `yaml:"destination" parser:"@@"`
}

type Add struct {
	Source *Param `yaml:"source" parser:"@@"`
}

type Sub struct {
	Source *Param `yaml:"source" parser:"@@"`
}

type Jmp struct {
	Label string `yaml:"label" parser:"@Ident"`
}

type Jez struct {
	Label string `yaml:"label" parser:"@Ident"`
}

type Jnz struct {
	Label string `yaml:"label" parser:"@Ident"`
}

type Jgz struct {
	Label string `yaml:"label" parser:"@Ident"`
}

type Jlz struct {
	Label string `yaml:"label" parser:"@Ident"`
}

type Jro struct {
	Source *Param `yaml:"source" parser:"@@"`
}

type Param struct {
	Register *Register `yaml:"register,omitempty" parser:"( @('ACC'|'BAK'|'NIL'|'LEFT'|'UP'|'RIGHT'|'DOWN'|'ANY'|'LAST') |"`
	Literal  *int      `yaml:"literal,omitempty" parser:"@Literal |"`
	Label    *string   `yaml:"label,omitempty" parser:"@Ident )"`
}
