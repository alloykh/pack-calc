package log

type NoopLogger struct {
}

func (n *NoopLogger) Trace(args ...any) {
}

func (n *NoopLogger) Info(args ...any) {
}

func (n *NoopLogger) Infof(template string, args ...any) {
}

func (n *NoopLogger) Warn(args ...any) {
}

func (n *NoopLogger) Error(args ...any) {
}

func (n *NoopLogger) Errorf(template string, args ...any) {
}

func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}
