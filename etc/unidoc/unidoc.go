package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/unidoc/unioffice/v2/spreadsheet"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	startTime := time.Now()
	ss := spreadsheet.New()
	sheet := ss.AddSheet()
	for range 102400 {
		row := sheet.AddRow()
		for range 50 {
			cell := row.AddCell()
			cell.SetString(RandStringBytes(6))
		}
	}
	if err := ss.Validate(); err != nil {
		fmt.Println(err)
		return
	}
	ss.SaveToFile("unidoc.xlsx")
	printBenchmarkInfo("unidoc", startTime)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(52)]
	}
	return string(b)
}

func printBenchmarkInfo(fn string, startTime time.Time) {
	var memStats runtime.MemStats
	var bToMb = func(b uint64) uint64 {
		return b / 1024 / 1024
	}
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Func: %s \tRSS = %v MB\tAlloc = %v MB\tTotalAlloc = %v MB\tSys = %v MB\tNumGC = %v \tCost = %s\n",
		fn, bToMb(maxRSS()), bToMb(memStats.Alloc), bToMb(memStats.TotalAlloc), bToMb(memStats.Sys), memStats.NumGC, time.Since(startTime))
}

// Func: unidoc 	RSS = 3275 MB	Alloc = 2650 MB	TotalAlloc = 5832 MB	Sys = 3285 MB	NumGC = 16 	Cost = 24.839927459s
