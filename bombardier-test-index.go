package main

import (
	"bytes"
	"fmt"
	"log"
	"os" // Импорт пакета os
	"os/exec"
	"strings"

	"github.com/olekukonko/tablewriter" // Импорт пакета tablewriter
)

func main() {
	// Команда и аргументы для выполнения
	cmd := exec.Command("bombardier", "-c", "1000", "-t", "60s", "-d", "60s", "-l", "http://192.168.22.92/index.html")

	// Массив для хранения результатов
	var results [][]string

	// Выполняем команду 10 раз
	for i := 1; i <= 1; i++ {
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Ошибка выполнения команды %d: %v, stderr: %s", i, err, stderr.String())
		}

		// Парсим вывод команды
		output := out.String()

		var requests, success, failed, avgLatency, maxLatency string
		for _, line := range strings.Split(output, "\n") {
			if strings.Contains(line, "Requests") {
				parts := strings.Fields(line)
				for i, part := range parts {
					if part == "Requests:" {
						requests = parts[i+1]
					} else if part == "Success:" {
						success = parts[i+1]
					} else if part == "Failed:" {
						failed = parts[i+1]
					} else if part == "Avg:" {
						avgLatency = parts[i+1]
					} else if part == "Max:" {
						maxLatency = parts[i+1]
					}
				}
			}
		}

		// Добавляем результаты в массив
		results = append(results, []string{
			fmt.Sprintf("%d", i),
			requests,
			success,
			failed,
			avgLatency,
			maxLatency,
		})
	}

	// Выводим результаты в таблицу
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Номер запроса", "Requests", "Success", "Failed", "Avg Latency", "Max Latency"}) // Исправленный вызов метода

	for _, row := range results {
		table.Append(row)
	}

	table.Render()
}
