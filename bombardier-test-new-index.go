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
	cmd := exec.Command("bombardier", "-c", "1000", "-t", "60s", "-d", "60s", "-l", "-p", "r", "http://192.168.22.92/index.html")

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

		// Парсинг статистики
		var AvgRps, StdevRps, MaxRps, reqsPerSec, latency, latency50, latency75, latency90, latency95, latency99, http1xx, http2xx, http3xx, http4xx, http5xx, otherHTTP, throughput string
		for _, line := range strings.Split(output, "\n") {
			if strings.Contains(line, "Avg") {
				parts := strings.Fields(line)
				AvgRps = parts[1]
				StdevRps = parts[2]
				MaxRps = parts[3]
			} else if strings.Contains(line, "Reqs/sec") {
				parts := strings.Fields(line)
				reqsPerSec = parts[1]
			} else if strings.Contains(line, "Latency") {
				parts := strings.Fields(line)
				latency = parts[1]
			} else if strings.Contains(line, "50%") {
				parts := strings.Fields(line)
				latency50 = parts[1]
			} else if strings.Contains(line, "75%") {
				parts := strings.Fields(line)
				latency75 = parts[1]
			} else if strings.Contains(line, "90%") {
				parts := strings.Fields(line)
				latency90 = parts[1]
			} else if strings.Contains(line, "95%") {
				parts := strings.Fields(line)
				latency95 = parts[1]
			} else if strings.Contains(line, "99%") {
				parts := strings.Fields(line)
				latency99 = parts[1]
			} else if strings.Contains(line, "HTTP codes") {
				http1xx = extractValue(line, "1xx")
				http2xx = extractValue(line, "2xx")
				http3xx = extractValue(line, "3xx")
				http4xx = extractValue(line, "4xx")
				http5xx = extractValue(line, "5xx")
				otherHTTP = extractValue(line, "others")
			} else if strings.Contains(line, "Throughput") {
				parts := strings.Fields(line)
				throughput = parts[1]
			}
		}

		// Добавляем результаты в массив
		results = append(results, []string{
			fmt.Sprintf("%d", i),
			reqsPerSec,
			latency,
			latency50,
			latency75,
			latency90,
			latency95,
			latency99,
			http1xx,
			http2xx,
			http3xx,
			http4xx,
			http5xx,
			otherHTTP,
			throughput,
		})
	}

	// Выводим результаты в таблицу
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Номер запроса", "Reqs/sec", "Latency", "50%", "75%", "90%", "95%", "99%", "1xx", "2xx", "3xx", "4xx", "5xx", "others", "Throughput"})
	for _, row := range results {
		table.Append(row)
	}
	table.Render()
}

func extractValue(line, key string) string {
	parts := strings.Split(line, key)
	if len(parts) > 1 {
		value := strings.Split(parts[1], ",")[0]
		return strings.TrimSpace(value)
	}
	return "0"
}

