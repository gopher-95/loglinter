package analyzer

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "проверка на правильность написания логов",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			fun, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			pkgId, ok := fun.X.(*ast.Ident)
			if !ok {
				return true
			}

			pkgName := pkgId.Name
			methodName := fun.Sel.Name

			if isLoggerCall(pkgName, methodName) {
				checkLogCall(pass, call)
			}

			return true
		})

	}

	return nil, nil
}

func isLoggerCall(pkgName, methodName string) bool {
	switch pkgName {
	case "log":
		return methodName == "Print" ||
			methodName == "Printf" ||
			methodName == "Println" ||
			methodName == "Fatal" ||
			methodName == "Fatalf" ||
			methodName == "Panic" ||
			methodName == "Panicf"
	case "slog":
		return methodName == "Debug" ||
			methodName == "Info" ||
			methodName == "Warn" ||
			methodName == "Error"
	case "zap":
		return methodName == "Debug" ||
			methodName == "Info" ||
			methodName == "Warn" ||
			methodName == "Error" ||
			methodName == "Fatal" ||
			methodName == "Panic"
	}
	return false
}

func checkLogCall(pass *analysis.Pass, call *ast.CallExpr) {
	for _, arg := range call.Args {
		lit, ok := arg.(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			continue
		}

		message := strings.Trim(lit.Value, "\"")

		checkFirstLetter(pass, lit.Pos(), message)

		checkEnglishOnly(pass, lit.Pos(), message)

		if checkSpecSymbolsAndEmojis(pass, lit.Pos(), message) {
			return
		}

		checkSensitiveData(pass, lit.Pos(), message)

		return
	}
}

func checkFirstLetter(pass *analysis.Pass, pos token.Pos, message string) {
	if message == "" {
		return
	}

	firstArg := []rune(message)[0]

	if unicode.IsLetter(firstArg) && unicode.IsUpper(firstArg) {
		pass.Reportf(pos, "лог должен начинаться с маленькой буквы")
	}
}

func checkEnglishOnly(pass *analysis.Pass, pos token.Pos, message string) {
	for _, r := range message {
		if unicode.IsLetter(r) {
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
				pass.Reportf(pos, "лог должен состоять только из английских букв")
				break
			}
		}
	}
}

func checkSpecSymbolsAndEmojis(pass *analysis.Pass, pos token.Pos, message string) bool {
	for _, r := range message {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			continue
		}

		pass.Reportf(pos, "лог не должен содержать спецсимволы или эмодзи")
		return true
	}

	return false
}

func checkSensitiveData(pass *analysis.Pass, pos token.Pos, message string) {
	lowerMsg := strings.ToLower(message)

	sensitiveWords := []string{
		"password",
		"pwd",
		"pass",
		"token",
		"api_key",
		"apikey",
		"api-key",
		"secret",
		"key",
		"auth",
		"credential",
		"jwt",
		"bearer",
		"authorization",
		"access_token",
		"refresh_token",
	}

	for _, word := range sensitiveWords {
		if strings.Contains(lowerMsg, word) {
			pass.Reportf(pos, "лог не должен содержать чувствительные данные")
			return
		}
	}
}
