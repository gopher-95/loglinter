package englishwords

import (
	"log/slog"

	"go.uber.org/zap"
)

func TestLogs() {
	//правильные логи - тест пропускаем
	//slog
	slog.Debug("debug message")
	slog.Info("info message")
	slog.Warn("warn message")
	slog.Error("error message")

	//zap
	logger := zap.NewExample()
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")

	//неправильные логи, тест не пропустит
	//slog
	slog.Debug("отладочное сообщение")    // want "лог должен состоять только из английских букв"
	slog.Info("информационное сообщение") // want "лог должен состоять только из английских букв"
	slog.Warn("предупреждение")           // want "лог должен состоять только из английских букв"
	slog.Error("ошибка соединения")       // want "лог должен состоять только из английских букв"

	//zap
	logger.Debug("отладка") // want "лог должен состоять только из английских букв"
	logger.Info("инфо")     // want "лог должен состоять только из английских букв"
	logger.Warn("внимание") // want "лог должен состоять только из английских букв"
	logger.Error("ошибка")  // want "лог должен состоять только из английских букв"

}
