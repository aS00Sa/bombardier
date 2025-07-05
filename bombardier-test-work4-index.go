package main
import (
	"bytes"
	"fmt"
	"log"
//	"os"
	"os/exec"
	"regexp"
	"strings"

//	"github.com/olekukonko/tablewriter"
)
func main() {
//	cmd := exec.Command("bombardier", "-c", "1000", "-t", "60s", "-d", "1s", "-l", "-p", "r", "http://192.168.22.92/index.html")
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
//        fmt.Println("i|requests|latency95|latency99|HTTP2xx|otherHTTP|Throughput")
	}
	// Выводим результаты в табличном формате с разделителями "|"
	fmt.Println("Номер запроса|Requests/s|Latency95%,ms|Latency99%,ms|HTTP 2xx|Others|Throughput,MB/s")
	for _, row := range results {
	        fmt.Println(strings.Join(row, "|"))
	}
}
