# bombardier [![Build Status](https://codesenberg.semaphoreci.com/badges/bombardier/branches/master.svg?key=249c678c-eb2a-441e-8128-1bdcfb9aaca6)](https://codesenberg.semaphoreci.com/projects/bombardier) [![Go Report Card](https://goreportcard.com/badge/github.com/codesenberg/bombardier)](https://goreportcard.com/report/github.com/codesenberg/bombardier) [![GoDoc](https://godoc.org/github.com/codesenberg/bombardier?status.svg)](http://godoc.org/github.com/codesenberg/bombardier)

bombardier is a HTTP(S) benchmarking tool. It is written in Go programming language and uses excellent [fasthttp](https://github.com/valyala/fasthttp) instead of Go's default http library, because of its lightning fast performance. 

With `bombardier v1.1` and higher you can now use `net/http` client if you need to test HTTP/2.x services or want to use a more RFC-compliant HTTP client.

## Installation
You can grab binaries in the [releases](https://github.com/codesenberg/bombardier/releases) section.
Alternatively, to get latest and greatest run:

Go 1.18+: `go install github.com/aS00Sa/bombardier@latest`

## Usage
```
bombardier [<flags>] <url>
```
For a more detailed information about flags consult [GoDoc](http://godoc.org/github.com/codesenberg/bombardier).

## Examples
Example of running `bombardier` against [this server](https://godoc.org/github.com/codesenberg/bombardier/cmd/utils/simplebenchserver):
```
root@debian1:/home/debian1/bombardier# bombardier --http1 -c 1000 -t 30s -d 600s -l -p r -o pt "http://192.168.22.92/proxy.php"
Statistics        Avg      Stdev        Max
  Reqs/sec      1475.18     909.21    5049.50
  Latency      618.73ms   196.54ms      1.65s
  Latency Distribution
     50%   614.29ms
     75%   658.74ms
     90%   783.24ms
     95%      1.02s
     99%      1.50s
  HTTP codes:
    1xx - 0, 2xx - 7404, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:   107.09MB/s
```
```
go get -u github.com/codesenberg/bombardier
go get -u github.com/olekukonko/tablewriter
```
```
# Edit CURL parameters
vi bombardier-test-index.go
vi bombardier-test-proxy.go
vi bombardier-test-redis.go

# ready steady go!
go run bombardier-test-index.go && go run bombardier-test-proxy.go && go run bombardier-test-redis.go & 

import index.html.txt proxy.php.txt redis.txt files to Calc
```



