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
        cmd := exec.Command("bombardier", "-c", "1000", "-t", "10s", "-d", "10s", "-l", "-p", "r", "http://192.168.22.92/index.html")
	// Массив для хранения результатов
	var results [][]string
	// Выполняем команду 1 раз
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
		var requests, latency95, latency99, HTTP2xx, otherHTTP, Throughput string
		for _, line := range strings.Split(output, "\n") {
			if strings.Contains(line, "Requests") {
				parts := strings.Fields(line)
				for i, part := range parts {
					if part == "Reqs/sec" {
						requests = parts[i+1]
					} else if part == "95%" {
						latency95 = parts[i+1]
					} else if part == "99%" {
						latency99 = parts[i+1]
					} else if part == "2xx" {
						HTTP2xx = parts[i+1]
					} else if part == "others" {
						otherHTTP = parts[i+1]
					} else if part == "Throughput:" {
						Throughput = parts[i+1]
					}
				}
			}
		}
	// Вывод статистики
	fmt.Println("\nStatistics")
	fmt.Printf("  Reqs/sec ", requests)
	fmt.Printf("    latency 95% ", latency95)
	fmt.Printf("    latency 99% ", latency99)
	fmt.Printf("    HTTP 2xx ", HTTP2xx)
	fmt.Printf("    others ", otherHTTP)
	fmt.Printf("    Throughput ", Throughput)
	fmt.Println("\n")
		// Добавляем результаты в массив
		results = append(results, []string{
			fmt.Sprintf("%d", i),
			requests,
			latency95,
			latency99,
			HTTP2xx,
			otherHTTP,
		})
	}
	// Выводим результаты в таблицу
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Номер запроса", "requests", "latency95", "latency99", "HTTP2xx", "otherHTTP"}) // Исправленный вызов метода
	for _, row := range results {
		table.Append(row)
	}
	table.Render()
}
