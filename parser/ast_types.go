package parser

type SimpleOp string

const (
	OpNop SimpleOp = "NOP"
	OpSwp SimpleOp = "SWP"
	OpSav SimpleOp = "SAV"
	OpNeg SimpleOp = "NEG"
)

type Register string

const (
	RegisterAcc   Register = "ACC"
	RegisterBak   Register = "BAK"
	RegisterNil   Register = "NIL"
	RegisterLeft  Register = "LEFT"
	RegisterUp    Register = "UP"
	RegisterRight Register = "RIGHT"
	RegisterDown  Register = "DOWN"
	RegisterAny   Register = "ANY"
	RegisterLast  Register = "LAST"
)
