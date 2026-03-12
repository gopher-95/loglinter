package sensitive

import (
	"log/slog"

	"go.uber.org/zap"
)

func TestLogs() {
	//правильные логи, тест должен пропустить
	//slog
	slog.Debug("request GET")
	slog.Info("operation 123")
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
	// password и вариации
	slog.Debug("user password")     // want "лог не должен содержать чувствительные данные"
	slog.Info("password is secret") // want "лог не должен содержать чувствительные данные"
	slog.Warn("invalid password")   // want "лог не должен содержать чувствительные данные"
	slog.Error("password expired")  // want "лог не должен содержать чувствительные данные"
	slog.Debug("my pwd is 123")     // want "лог не должен содержать чувствительные данные"
	slog.Info("pwd secret")         // want "лог не должен содержать чувствительные данные"
	slog.Warn("pass 12345")         // want "лог не должен содержать чувствительные данные"

	// token и вариации
	slog.Debug("using token")             // want "лог не должен содержать чувствительные данные"
	slog.Info("token expired")            // want "лог не должен содержать чувствительные данные"
	slog.Warn("invalid token")            // want "лог не должен содержать чувствительные данные"
	slog.Error("token validation failed") // want "лог не должен содержать чувствительные данные"
	slog.Debug("accesstoken abc")         // want "лог не должен содержать чувствительные данные"
	slog.Info("refreshtoken updated")     // want "лог не должен содержать чувствительные данные"

	// api key и вариации
	slog.Debug("api key is 123") // want "лог не должен содержать чувствительные данные"
	slog.Info("apikey abc")      // want "лог не должен содержать чувствительные данные"
	slog.Warn("api key xyz")     // want "лог не должен содержать чувствительные данные"
	slog.Error("apikey invalid") // want "лог не должен содержать чувствительные данные"

	// secret и key
	slog.Debug("secret value") // want "лог не должен содержать чувствительные данные"
	slog.Info("secret data")   // want "лог не должен содержать чувствительные данные"
	slog.Warn("key not found") // want "лог не должен содержать чувствительные данные"
	slog.Error("invalid key")  // want "лог не должен содержать чувствительные данные"

	// auth и вариации
	slog.Debug("auth required")       // want "лог не должен содержать чувствительные данные"
	slog.Info("authorization failed") // want "лог не должен содержать чувствительные данные"
	slog.Warn("unauthorized access")  // want "лог не должен содержать чувствительные данные"

	// jwt и bearer
	slog.Debug("jwt validation")        // want "лог не должен содержать чувствительные данные"
	slog.Info("jwt token expired")      // want "лог не должен содержать чувствительные данные"
	slog.Warn("bearer token invalid")   // want "лог не должен содержать чувствительные данные"
	slog.Error("invalid jwt signature") // want "лог не должен содержать чувствительные данные"

	// credential
	slog.Debug("invalid credentials") // want "лог не должен содержать чувствительные данные"
	slog.Info("credential error")     // want "лог не должен содержать чувствительные данные"

	// В составе слов
	slog.Debug("passphrase entered") // want "лог не должен содержать чувствительные данные"
	slog.Info("keychain unlocked")   // want "лог не должен содержать чувствительные данные"
	slog.Warn("authservice timeout") // want "лог не должен содержать чувствительные данные"
	slog.Error("tokenization error") // want "лог не должен содержать чувствительные данные"

	//zap
	// password и вариации
	logger.Debug("user password")     // want "лог не должен содержать чувствительные данные"
	logger.Info("password is secret") // want "лог не должен содержать чувствительные данные"
	logger.Warn("invalid password")   // want "лог не должен содержать чувствительные данные"
	logger.Error("password expired")  // want "лог не должен содержать чувствительные данные"
	logger.Debug("my pwd is 123")     // want "лог не должен содержать чувствительные данные"
	logger.Info("pwd secret")         // want "лог не должен содержать чувствительные данные"

	// token и вариации
	logger.Debug("using token")             // want "лог не должен содержать чувствительные данные"
	logger.Info("token expired")            // want "лог не должен содержать чувствительные данные"
	logger.Warn("invalid token")            // want "лог не должен содержать чувствительные данные"
	logger.Error("token validation failed") // want "лог не должен содержать чувствительные данные"
	logger.Debug("accesstoken abc")         // want "лог не должен содержать чувствительные данные"
	logger.Info("refreshtoken updated")     // want "лог не должен содержать чувствительные данные"

	// api key и вариации
	logger.Debug("api key is 123") // want "лог не должен содержать чувствительные данные"
	logger.Info("apikey abc")      // want "лог не должен содержать чувствительные данные"
	logger.Warn("apikey xyz")      // want "лог не должен содержать чувствительные данные"
	logger.Error("apikey invalid") // want "лог не должен содержать чувствительные данные"

	// secret и key
	logger.Debug("secret value") // want "лог не должен содержать чувствительные данные"
	logger.Info("secret data")   // want "лог не должен содержать чувствительные данные"
	logger.Warn("key not found") // want "лог не должен содержать чувствительные данные"
	logger.Error("invalid key")  // want "лог не должен содержать чувствительные данные"

	// auth и вариации
	logger.Debug("auth required")       // want "лог не должен содержать чувствительные данные"
	logger.Info("authorization failed") // want "лог не должен содержать чувствительные данные"
	logger.Warn("unauthorized access")  // want "лог не должен содержать чувствительные данные"

	// jwt и bearer
	logger.Debug("jwt validation")        // want "лог не должен содержать чувствительные данные"
	logger.Info("jwt token expired")      // want "лог не должен содержать чувствительные данные"
	logger.Warn("bearer token invalid")   // want "лог не должен содержать чувствительные данные"
	logger.Error("invalid jwt signature") // want "лог не должен содержать чувствительные данные"
	logger.Debug("invalid credentials")   // want "лог не должен содержать чувствительные данные"
	logger.Info("credential error")       // want "лог не должен содержать чувствительные данные"

	// В составе слов
	logger.Debug("passphrase entered") // want "лог не должен содержать чувствительные данные"
	logger.Info("keychain unlocked")   // want "лог не должен содержать чувствительные данные"
	logger.Warn("authservice timeout") // want "лог не должен содержать чувствительные данные"
	logger.Error("tokenization error") // want "лог не должен содержать чувствительные данные"

}
