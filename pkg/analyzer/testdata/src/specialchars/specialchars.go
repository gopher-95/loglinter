package specialchars

import (
	"log/slog"

	"go.uber.org/zap"
)

func TestLogs() {
	//правильные логи, тест пропускает
	//slog
	slog.Debug("debug message")
	slog.Info("info message")
	slog.Warn("warn message")
	slog.Error("error message")
	slog.Debug("user logged in")
	slog.Info("2 users connected")
	slog.Warn("connection timeout")
	slog.Error("database error 404")
	slog.Debug("")

	//zap
	logger := zap.NewExample()
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
	logger.Debug("user logged in")
	logger.Info("2 users connected")
	logger.Warn("connection timeout")
	logger.Error("database error 404")

	//неправильные логи, тест не пропустит
	//slog
	//эмодзи
	slog.Info("server started 🚀")           // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("error 🔥 connection failed") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("hello 😊 world")             // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("warning ⚠️")                 // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Info("crash 💔")                    // want "лог не должен содержать спецсимволы или эмодзи"

	//восклицательные знаки
	slog.Info("connection failed!") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("error!!")           // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("warning!!!")         // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("urgent!")           // want "лог не должен содержать спецсимволы или эмодзи"

	//точки и многоточия
	slog.Info("downloading.")  // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("processing..") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("waiting...")    // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("timeout....")  // want "лог не должен содержать спецсимволы или эмодзи"

	//вопросительные знаки
	slog.Info("are you sure?") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("what??")       // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("really???")     // want "лог не должен содержать спецсимволы или эмодзи"

	//двоеточия и точки с запятой
	slog.Info("warning: low disk") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("error: database")  // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("status: ok")       // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("hello; world")      // want "лог не должен содержать спецсимволы или эмодзи"

	//запятые и дефисы
	slog.Info("hello, world")        // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("timeout - retrying") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("one,two,three")       // want "лог не должен содержать спецсимволы или эмодзи"

	//спецсимволы
	slog.Info("user@email.com")  // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("price $100")     // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("error #123")      // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("10% done")       // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Info("key^value")       // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("a & b")          // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("test*test")       // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("path/to/file")   // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Info("key=value")       // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("pipe|separator") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("back\\slash")     // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("`quote`")        // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Info("~tilde~")         // want "лог не должен содержать спецсимволы или эмодзи"

	//скобки и кавычки
	slog.Info("(parentheses)") // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Debug("[brackets]")   // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Warn("{braces}")      // want "лог не должен содержать спецсимволы или эмодзи"
	slog.Error("\"quotes\"")   // want "лог не должен содержать спецсимволы или эмодзи"

	//zap
	//эмодзи
	logger.Info("server started 🚀")           // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Error("error 🔥 connection failed") // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Debug("hello 😊 world")             // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Warn("warning ⚠️")                 // want "лог не должен содержать спецсимволы или эмодзи"

	//восклицательные знаки
	logger.Info("connection failed!") // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Error("error!!")           // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Warn("warning!!!")         // want "лог не должен содержать спецсимволы или эмодзи"

	//точки и многоточия
	logger.Info("downloading.")  // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Debug("processing..") // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Warn("waiting...")    // want "лог не должен содержать спецсимволы или эмодзи"

	//вопросительные знаки
	logger.Info("are you sure?") // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Debug("what??")       // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Warn("really???")     // want "лог не должен содержать спецсимволы или эмодзи"

	//двоеточия и точки с запятой
	logger.Info("warning: low disk") // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Error("error: database")  // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Warn("hello; world")      // want "лог не должен содержать спецсимволы или эмодзи"

	//запятые и дефисы
	logger.Info("hello, world")        // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Debug("timeout - retrying") // want "лог не должен содержать спецсимволы или эмодзи"

	//спецсимволы
	logger.Info("user@email.com") // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Debug("price $100")    // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Warn("error #123")     // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Error("10% done")      // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Info("key^value")      // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Debug("a & b")         // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Warn("test*test")      // want "лог не должен содержать спецсимволы или эмодзи"
	logger.Error("path/to/file")  // want "лог не должен содержать спецсимволы или эмодзи"
}
