// createDefaultStruct.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Структура папок
	dirs := []string{
		"cmd/app",
		"internal/api",
		"internal/config",
		"pkg/utils",
	}

	// Создание папок
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Ошибка при создании директории %s: %v\n", dir, err)
			return
		}
		fmt.Printf("Создана директория: %s\n", dir)
	}

	// Создание файла .golangci.yml с содержимым
	ymlContent := `
linters:
  enable-all: true  # Включает все линтеры
  disable:
    - gomnd # вроде то же что и "mnd"
    - forbidigo # что-то сложное
    - exhaustruct # все поля должны быть проинициализированны
    - mnd # Детектит все числа, которые не const
    #- godox
    - exhaustivestruct # вроде то же что и "exhaustruct"
    - depguard # Душка с импортами
run:
  timeout: 20m        # Устанавливает таймаут для выполнения линтинга
issues:
  max-issues-per-linter: 0  # Неограниченное количество проблем на один линтер
  max-same-issues: 0        # Неограниченное количество одинаковых проблем
  exclude-use-default: false
linters-settings:
  govet:
    enable-all: true
  staticcheck:
    checks: ["all"]
  gosimple:
    checks: ["all"]
  gosec:
    excludes:
      - G404  # Отключение правила G404 Insecure random number source (rand)
  wsl:
    allow-cuddle-declarations: true
  varnamelen:
    max-distance: 5
    min-name-length: 3
    check-receiver: false
    check-return: false
    check-type-param: false
    ignore-type-assert-ok: false
    ignore-map-index-ok: false
    ignore-chan-recv-ok: false
    ignore-decls:
      - w http.ResponseWriter
      - r *http.Request
      - db *sql.DB
      - wg *sync.WaitGroup
      - mu *sync.Mutex
      - ok bool
  #cyclop:
  # package-average: 0.1
  gocritic: # вроде чекает паттерны
    enable-all: true
  gocyclo:
    min-complexity: 10
`
	ymlPath := filepath.Join(".", ".golangci.yml")
	if err := os.WriteFile(ymlPath, []byte(ymlContent), 0644); err != nil {
		fmt.Printf("Ошибка при создании файла .golangci.yml: %v\n", err)
		return
	}
	fmt.Println("Файл .golangci.yml успешно создан.")

	gitIgnore := `
.git*
.vsc*
.env*
staticcheck.conf
.golangci.yml
go.sum
*workspace*
`

	gitIgnorePath := filepath.Join(".", ".gitignore")
	if err := os.WriteFile(gitIgnorePath, []byte(gitIgnore), 0644); err != nil {
		fmt.Printf("Ошибка при создании файла .gitignore: %v\n", err)
		return
	}
}
