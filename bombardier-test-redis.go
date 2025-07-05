package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)


// Определяем структуру для хранения параметров
type BombardierParams struct {
	URL        string
	Connections string
	Timeout    string
	Duration   string
}

func main() {
	// Инициализируем параметры
	params := BombardierParams{
		URL:        "http://192.168.22.92/redis.php",
		Connections: "1000",
		Timeout:    "30s",
		Duration:   "60s",
	}

	var results [][]string
	for i := 1; i <= 10; i++ {
		// Создаем команду с параметрами
		cmd := exec.Command(
			"bombardier",
			"-c", params.Connections,
			"-t", params.Timeout,
			"-d", params.Duration,
			"-l",
			"-p", "r",
			params.URL,
		)

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
                var unitLatency95, unitLatency99, unitThroughput string
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
                        unitLatency95 = latency95Match[2]
//			if len(latency95Match) > 2 && latency95Match[2] != "" {
//				latency95 += latency95Match[2]
//			} 
//                        else if len(latency95Match) = 1 && latency95Match[2] != "" {
//				latency95 += "ms"
//			}
		} else {
			latency95 = "0"
                        unitLatency95 = ""
		}

		// Извлечение Latency 99%
		latency99Match := regexLatency99.FindStringSubmatch(output)
		if len(latency99Match) > 1 {
			latency99 = latency99Match[1]
                        unitLatency99 = latency99Match[2]
//			if len(latency99Match) > 2 && latency99Match[2] != "" {
//				latency99 += latency99Match[2]
//			} 
//                        else if len(latency99Match) = 1 && latency99Match[2] != "" {
//				latency99 += "ms"
//			}
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
                        unitThroughput = throughputMatch[2]
//			if len(throughputMatch) > 2 && throughputMatch[2] != "" {
//				Throughput += throughputMatch[2]
//			} 
//                        else {
//				Throughput += "MB/s"
//			}
		} else {
			Throughput = "0"
		}

		// Добавляем результаты в массив
		results = append(results, []string{
			fmt.Sprintf("%d", i),
                        params.URL,
                        params.Connections,
                        params.Duration,
                        params.Timeout,
                        requests,
			latency95,
                        unitLatency95,
			latency99,
                        unitLatency99,
			HTTP2xx,
			otherHTTP,
			Throughput,
                        unitThroughput,
		})
		// Выводим промежуточные сырые данные каждого цикла замеров
//		fmt.Printf("index.html %d, requests: %s, latency95: %s, latency99: %s, HTTP2xx: %s, otherHTTP: %s, Throughput: %s\n", i, requests, latency95, latency99, HTTP2xx, otherHTTP, Throughput)
	}

	// Выводим результаты в табличном формате с разделителями "|"
	fmt.Println("Count|URL|Connections|Duration|Timeout|Requests/s|Latency95%|unit|Latency99%|unit|HTTP2xx|OthersHTTP|Throughput|unit")
	for _, row := range results {
		fmt.Println(strings.Join(row, "|"))
	}
}

