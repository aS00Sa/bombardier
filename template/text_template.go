package main

import (
	"os"
	"text/template"
)

func main() {
	result := Result{
		Target:      "http://debian2/proxy.php",
		Duration:    "60s",
		Connections: 1000,
		ProgressBar: "===================================================================================================================",
		RequestsPerSec: struct {
			Avg   float64
			Stdev float64
			Max   float64
		}{
			Avg:   1347.49,
			Stdev: 350.99,
			Max:   2456.43,
		},
		Latency: struct {
			Avg   float64
			Stdev float64
			Max   float64
			P50   float64
			P75   float64
			P90   float64
			P95   float64
			P99   float64
		}{
			Avg:   710.07,
			Stdev: 527.81,
			Max:   8.86,
			P50:   616.11,
			P75:   823.18,
			P90:   1.26,
			P95:   1.64,
			P99:   2.85,
		},
		HTTPCodes: struct {
			Code1xx int
			Code2xx int
			Code3xx int
			Code4xx int
			Code5xx int
			Others  int
		}{
			Code1xx: 0,
			Code2xx: 14452,
			Code3xx: 0,
			Code4xx: 0,
			Code5xx: 0,
			Others:  0,
		},
		Throughput: 84.56,
	}

	tmpl := template.Must(template.New("result").Parse(resultTemplate))
	err := tmpl.Execute(os.Stdout, result)
	if err != nil {
		panic(err)
	}
}

type Result struct {
    Target      string
    Duration    string
    Connections int
    ProgressBar string
    
    RequestsPerSec struct {
        Avg   float64
        Stdev float64
        Max   float64
    }
    
    Latency struct {
        Avg   float64 // в миллисекундах
        Stdev float64 // в миллисекундах
        Max   float64 // в секундах
        P50   float64 // в миллисекундах
        P75   float64 // в миллисекундах
        P90   float64 // в секундах
        P95   float64 // в секундах
        P99   float64 // в секундах
    }
    
    HTTPCodes struct {
        Code1xx int
        Code2xx int
        Code3xx int
        Code4xx int
        Code5xx int
        Others  int
    }
    
    Throughput float64 // в MB/s
}

const (
	plainTextTemplate = `
{{define "resultTemplate"}}
Bombarding {{.Target}} for {{.Duration}} using {{.Connections}} connection(s)
[{{.ProgressBar}}] {{.Duration}}
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec      {{.RequestsPerSec.Avg | printf "%.2f"}}    {{.RequestsPerSec.Stdev | printf "%.2f"}}   {{.RequestsPerSec.Max | printf "%.2f"}}
  Latency      {{.Latency.Avg | printf "%.2fms"}}   {{.Latency.Stdev | printf "%.2fms"}}  {{.Latency.Max | printf "%.2fs"}}
  Latency Distribution
     50%   {{.Latency.P50 | printf "%.2fms"}}
     75%   {{.Latency.P75 | printf "%.2fms"}}
     90%   {{.Latency.P90 | printf "%.2fs"}}
     95%   {{.Latency.P95 | printf "%.2fs"}}
     99%   {{.Latency.P99 | printf "%.2fs"}}
  HTTP codes:
    1xx - {{.HTTPCodes.Code1xx}}, 2xx - {{.HTTPCodes.Code2xx}}, 3xx - {{.HTTPCodes.Code3xx}}, 4xx - {{.HTTPCodes.Code4xx}}, 5xx - {{.HTTPCodes.Code5xx}}
    others - {{.HTTPCodes.Others}}
  Throughput:    {{.Throughput | printf "%.2fMB/s"}}
{{- end -}}`
)

