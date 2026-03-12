package zap

// Logger представляет zap логгер
type Logger struct{}

func NewExample() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(msg string, fields ...any) {}

func (l *Logger) Info(msg string, fields ...any) {}

func (l *Logger) Warn(msg string, fields ...any) {}

func (l *Logger) Error(msg string, fields ...any) {}
