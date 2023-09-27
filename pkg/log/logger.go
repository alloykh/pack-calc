package log

type Logger interface {
	Trace(...any)
	Info(...any)
	Infof(string, ...any)
	Warn(...any)
	Error(...any)
	Errorf(string, ...any)
}
