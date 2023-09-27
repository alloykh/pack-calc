package log

//go:generate enumer -type=Verbosity -json -text -yaml -transform=snake
type Verbosity int

const (
	Silent Verbosity = iota
	Error
	Info
	Warning
	Trace
)
