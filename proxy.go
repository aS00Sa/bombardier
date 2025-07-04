package main

import (
	"os"
	"text/template"
)

func main() {
	result := Result{
		Target:      "http://debian2/proxy.php",
		Duration:    "10s",
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

