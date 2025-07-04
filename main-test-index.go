package main

import (
        "fmt"
        "log"
        "os"
        "time"

        "github.com/codesenberg/bombardier"
        "github.com/olekukonko/tablewriter"
)

// Структура для хранения результатов запроса
type RequestResult struct {
        StatusCode int
        Duration   time.Duration
}

func main() {
        // URL, на который будем отправлять запросы
        url := "http://192.168.22.92/index.html"

        // Создаем клиент Bombardier
        client := bombardier.NewClient()

        // Массив для хранения результатов
        results := make([]RequestResult, 10)

        // Выполняем 10 запросов
        for i := 0; i < 10; i++ {
                start := time.Now()
                resp, err := client.Get(url)
                if err != nil {
                        log.Fatalf("Ошибка выполнения запроса: %v", err)
                }
                duration := time.Since(start)
                results[i] = RequestResult{
                        StatusCode: resp.StatusCode,
                        Duration:   duration,
                }
        }

        // Выводим результаты в виде таблицы
        table := tablewriter.NewWriter(os.Stdout)
        table.SetHeader([]string{"Номер запроса", "Статус код", "Длительность (ms)"})

        for i, res := range results {
                table.Append([]string{
                        fmt.Sprintf("%d", i+1),
                        fmt.Sprintf("%d", res.StatusCode),
                        fmt.Sprintf("%.2f", res.Duration.Seconds()*1000),
                })
        }

        table.Render()
}

