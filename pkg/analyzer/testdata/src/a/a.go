package a

import (
	"log"
	"log/slog"
)

func TestLogs() {
	log.Print("starting server")
	slog.Info("user logged in")
	log.Print("2 users connected")
	log.Print("connection timeout retrying")
	log.Print("server started")
	log.Print("are you sure")
	log.Print("warning low disk")
	log.Print("checking connection please wait")

	// Неправильные логи (эмодзи)
	log.Print("server started 🚀")          // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("error 🔥 connection failed") // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("hello 😊 world")             // want "лог не должен содержать спецсимволы или эмодзи"

	// Неправильные логи (любые знаки препинания)
	log.Print("connection failed!")   // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("connection failed!!")  // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("connection failed!!!") // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("downloading.")         // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("downloading..")        // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("downloading...")       // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("what?")                // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("what??")               // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("what???")              // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("warning: low disk")    // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("timeout - retrying")   // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("hello, world")         // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("hello; world")         // want "лог не должен содержать спецсимволы или эмодзи"

	// Неправильные логи (спецсимволы)
	log.Print("user@email.com") // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("price $100")     // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("error #123")     // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("10% done")       // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("key^value")      // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("a & b")          // want "лог не должен содержать спецсимволы или эмодзи"
	log.Print("test*test")      // want "лог не должен содержать спецсимволы или эмодзи"

	log.Print("user password")      // want "лог не должен содержать чувствительные данные"
	log.Print("password is secret") // want "лог не должен содержать чувствительные данные"
	log.Print("my pwd is 123")      // want "лог не должен содержать чувствительные данные"
	log.Print("using token")        // want "лог не должен содержать чувствительные данные"
	log.Print("api key is 123")     // want "лог не должен содержать чувствительные данные"
	log.Print("secret value")       // want "лог не должен содержать чувствительные данные"
	log.Print("auth required")      // want "лог не должен содержать чувствительные данные"
	log.Print("jwt validation")     // want "лог не должен содержать чувствительные данные"
}
