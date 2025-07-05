package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	var results [][]string
	for i := 1; i <= 10; i++ {
		cmd := exec.Command("bombardier", "-c", "1000", "-t", "60s", "-d", "1s", "-l", "-p", "r", "http://192.168.22.92/index.html")
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
		regexLatency95 := regexp.MustCompile(`95%\s+([\d\.]+)(\w+)?`)
		regexLatency99 := regexp.MustCompile(`99%\s+([\d\.]+)(\w+)?`)
		regexHTTP2xx := regexp.MustCompile(`2xx -\s+([\d\.]+)`)
		regexOthers := regexp.MustCompile(`others -\s+([\d\.]+)`)
		regexThroughput := regexp.MustCompile(`Throughput:\s+([\d\.]+)(\w+)?`)

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
			if len(latency95Match) > 2 && latency95Match[2] != "" {
				latency95 += latency95Match[2]
			} else {
				latency95 += "ms"
			}
		} else {
			latency95 = "0"
		}

		// Извлечение Latency 99%
		latency99Match := regexLatency99.FindStringSubmatch(output)
		if len(latency99Match) > 1 {
			latency99 = latency99Match[1]
			if len(latency99Match) > 2 && latency99Match[2] != "" {
				latency99 += latency99Match[2]
			} else {
				latency99 += "ms"
			}
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
			if len(throughputMatch) > 2 && throughputMatch[2] != "" {
				Throughput += throughputMatch[2]
			} else {
				Throughput += "MB/s"
			}
		} else {
			Throughput = "0"
		}

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
		// Выводим промежуточные сырые данные каждого цикла замеров
		fmt.Printf("index.html %d, requests: %s, latency95: %s, latency99: %s, HTTP2xx: %s, otherHTTP: %s, Throughput: %s\n", i, requests, latency95, latency99, HTTP2xx, otherHTTP, Throughput)
	}

	// Выводим результаты в табличном формате с разделителями "|"
	fmt.Println("Номер запроса|Requests/s|Latency95%|Latency99%|HTTP 2xx|Others|Throughput")
	for _, row := range results {
		fmt.Println(strings.Join(row, "|"))
	}
}

