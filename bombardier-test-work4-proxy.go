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
	cmd := exec.Command("bombardier", "-c", "1000", "-t", "60s", "-d", "60s", "-l", "-p", "r", "http://192.168.22.92/proxy.php")
	var results [][]string
	for i := 1; i <= 10; i++ {
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
		regexLatency95 := regexp.MustCompile(`95%\s+([\d\.]+)ms`)
		regexLatency99 := regexp.MustCompile(`99%\s+([\d\.]+)ms`)
		regexHTTP2xx := regexp.MustCompile(`2xx -\s+([\d\.]+)`)
		regexOthers := regexp.MustCompile(`others -\s+([\d\.]+)`)
		regexThroughput := regexp.MustCompile(`Throughput:\s+([\d\.]+)MB/s`)

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
			latency95 = latency95Match[1] + "ms"
		} else {
			latency95 = "0ms"
		}

		// Извлечение Latency 99%
		latency99Match := regexLatency99.FindStringSubmatch(output)
		if len(latency99Match) > 1 {
			latency99 = latency99Match[1] + "ms"
		} else {
			latency99 = "0ms"
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
			Throughput = throughputMatch[1] + "MB/s"
		} else {
			Throughput = "0MB/s"
		}

		// Вывод статистики
		fmt.Println("\nStatistics")
		fmt.Printf("  Reqs/sec %s\n", requests)
		fmt.Printf("    latency 95% %s\n", latency95)
		fmt.Printf("    latency 99% %s\n", latency99)
		fmt.Printf("    HTTP 2xx %s\n", HTTP2xx)
		fmt.Printf("    others %s\n", otherHTTP)
		fmt.Printf("    Throughput %s\n", Throughput)
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

