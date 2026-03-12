package firstletter

import (
	"log/slog"

	"go.uber.org/zap"
)

func TestLogs() {
	//правильные логи, которые тест должен пропустить
	// slog
	slog.Debug("debug message")
	slog.Info("info message")
	slog.Warn("warning message")
	slog.Error("error message")

	// zap
	logger := zap.NewExample()
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	//неправильные логи, тест не пропустит
	// slog
	slog.Debug("Debug message") // want "лог должен начинаться с маленькой буквы"
	slog.Info("Info message")   // want "лог должен начинаться с маленькой буквы"
	slog.Warn("Warn message")   // want "лог должен начинаться с маленькой буквы"
	slog.Error("Error message") // want "лог должен начинаться с маленькой буквы"

	//zap
	logger.Debug("Debug message") // want "лог должен начинаться с маленькой буквы"
	logger.Info("Info message")   // want "лог должен начинаться с маленькой буквы"
	logger.Warn("Warn message")   // want "лог должен начинаться с маленькой буквы"
	logger.Error("Error message") // want "лог должен начинаться с маленькой буквы"
}
