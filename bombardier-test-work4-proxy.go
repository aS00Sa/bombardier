package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
//	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	cmd := exec.Command("bombardier", "-c", "1000", "-t", "60s", "-d", "1s", "-l", "-p", "r", "http://192.168.22.92/proxy.php")
	var results [][]string
	for i := 1; i <= 1; i++ {
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Ошибка выполнения команды %d: %v, stderr: %s", i, err, stderr.String())
		}

		output := out.String()
		var requests, latency95, latency99, HTTP2xx, otherHTTP, Throughput string

		// Регулярные выражения для парсинга
		regexReqsPerSec := regexp.MustCompile(`Reqs/sec\s+([\d\.]+)`)
		regexLatency95 := regexp.MustCompile(`95%\s+([\d\.]+)`)
		regexLatency99 := regexp.MustCompile(`99%\s+([\d\.]+)`)
		regexHTTP2xx := regexp.MustCompile(`2xx -\s+([\d\.]+)`)
		regexOthers := regexp.MustCompile(`others -\s+([\d\.]+)`)
		regexThroughput := regexp.MustCompile(`Throughput:\s+([\d\.]+)`)

		// Извлечение Reqs/sec
		reqsPerSecMatch := regexReqsPerSec.FindStringSubmatch(output)
		if len(reqsPerSecMatch) > 1 {
			requests = reqsPerSecMatch[1]
		} else {
			requests = "0"
		}

		// Извлечение Latency 95%
		latency95Match := regexLatency95.FindStringSubmatch(output)
		if len(latency95Match) > 1 {
			latency95 = latency95Match[1]
		} else {
			latency95 = "0"
		}

		// Извлечение Latency 99%
		latency99Match := regexLatency99.FindStringSubmatch(output)
		if len(latency99Match) > 1 {
			latency99 = latency99Match[1]
		} else {
			latency99 = "0"
		}

		// Извлечение HTTP 2xx
		http2xxMatch := regexHTTP2xx.FindStringSubmatch(output)
		if len(http2xxMatch) > 1 {
			HTTP2xx = http2xxMatch[1]
		} else {
			HTTP2xx = "0"
		}

		// Извлечение others
		othersMatch := regexOthers.FindStringSubmatch(output)
		if len(othersMatch) > 1 {
			otherHTTP = othersMatch[1]
		} else {
			otherHTTP = "0"
		}

		// Извлечение Throughput
		throughputMatch := regexThroughput.FindStringSubmatch(output)
		if len(throughputMatch) > 1 {
			Throughput = throughputMatch[1]
		} else {
			Throughput = "0"
		}

		// Вывод статистики
		fmt.Println("\nStatistics")
		fmt.Printf("\n  Reqs/sec %s", requests)
		fmt.Printf("\n    latency 95% %s", latency95)
		fmt.Printf("\n    latency 99% %s", latency99)
		fmt.Printf("\n    HTTP 2xx %s", HTTP2xx)
		fmt.Printf("\n    others %s", otherHTTP)
		fmt.Printf("\n    Throughput %s", Throughput)
		fmt.Println("\n")

		// Добавляем результаты в массив
		results = append(results, []string{
			fmt.Sprintf("%d", i),
			requests,
			latency95,
			latency99,
			HTTP2xx,
			otherHTTP,
			Throughput,
		})
	}

	// Выводим результаты в таблицу
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Номер запроса", "Requests", "Latency 95%", "Latency 99%", "HTTP 2xx", "Others", "Throughput"})
	for _, row := range results {
		table.Append(row)
	}
	table.Render()
}

