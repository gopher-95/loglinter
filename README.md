# loglinter

Линтер для проверки лог-записей в Go-проектах. Анализирует код и проверяет соответствие логов четырем правилам.

### О проекте

Проект представляет собой статический анализатор кода для Go, который находит в коде вызовы логгеров (`slog`, `zap`) и проверяет их аргументы на соответствие правилам форматирования и безопасности.

### Технологии
- **Язык:** Go 1.25+
- **Основной пакет:** `golang.org/x/tools/go/analysis`
- **Поддерживаемые логгеры:** `log/slog`, `go.uber.org/zap`
- **Тестирование:** встроенный пакет `testing` + `analysistest`

### Структура проекта
```
loglinter/
├── pkg/
│   └── analyzer/
│       ├── analyzer.go          # Основная логика линтера (4 правила)
│       ├── analyzer_test.go     # Unit-тесты
│       └── testdata/
│           └── src/
│               ├── firstletter/     # Тесты для правила 1
│               │   └── firstletter.go
│               ├── englishwords/    # Тесты для правила 2
│               │   └── englishwords.go
│               ├── specialchars/    # Тесты для правила 3
│               │   └── specialchars.go
│               ├── sensitive/       # Тесты для правила 4
│               │   └── sensitive.go
│               └── go.uber.org/
│                   └── zap/         # Заглушка для zap
│                       └── zap.go
│
├── plugin/
│   └── plugin.go                  # Плагин для golangci-lint
│
├── cmd/
│   └── loglinter/
│       └── main.go                # Точка входа для go vet
│
├── Dockerfile.demo                # Docker-образ для кросс-платформенного использования
├── .golangci.yml                  # Конфигурация для плагина
├── go.mod                         # Модуль и зависимости
└── README.md                      # Документация
```

# Unit тесты

Все тесты расположены в директории `pkg/analyzer/testdata/src/`. Для каждого правила создана отдельная папка с тестовыми файлами:
- `firstletter/` - проверка на строчную букву
- `englishwords/` - проверка на английский язык
- `specialchars/` - проверка на спецсимволы и эмодзи
- `sensitive/` - проверка на чувствительные данные

## Запуск тестов из директории пакета

1. Перейдите в директорию с тестами:
```bash
cd pkg/analyzer
```
2. Выполните команду 
```bash
go test -v
```


# Сборка плагина (только для Linux/macOS)
1. Клонировать репозиторий 
```bash
git clone https://github.com/gopher-95/loglinter.git
```
2. Собрать плагин
```bash
go build -buildmode=plugin -o loglinter.so plugin/plugin.go
```

## Использование в проекте
1. Скопируйте плагин в корень вашего проекта:
```bash
cp loglinter.so /path/to/your/project/
```
2. Создайте файл .golangci.yml в корне проекта:
```yaml
linters-settings:
  custom:
    loglinter:
      path: ./loglinter.so
      description: "Линтер для проверки логов"
linters:
  enable:
    - loglinter
```
3. Запустите линтер
```bash
golangci-lint run
```

# Проблема с Windows 
На платформе Windows Go не поддерживает сборку и загрузку плагинов (-buildmode=plugin). Это техническое ограничение языка Go, которое проявляется ошибкой:
```bash
-buildmode=plugin not supported on windows/amd64
```
В качестве решения этой проблемы было принято создать Dockerfile в корне проекта.
Для использования на Windows необходимо склонировать мой репозиторий
```bash
git clone https://github.com/gopher-95/loglinter.git
```
Перейти в него
```bash
cd loglinter
```
Собрать образ локально
```bash
docker build -f Dockerfile.demo -t loglinter-local .
```
Перейти в свой проект
```bash
cd /path/to/their-project
```
Использовать в своем проекте
```bash
docker run --rm -v ${PWD}:/app -w /app loglinter-demo sh -c "cd /app && go vet -vettool=/go/bin/loglinter ./..."
```
# Я бы хотел продемонстрировать работу своего линтера на своем проекте
### В корне проекта будет лежать `.pdf` файл, в котором я запустил свой линтер (последним способом с помощью Docker)


