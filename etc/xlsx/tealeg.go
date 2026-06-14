package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/tealeg/xlsx/v3"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	startTime := time.Now()
	f := xlsx.NewFile()
	sheet, err := f.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for range 102400 {
		row := sheet.AddRow()
		for range 50 {
			cell := row.AddCell()
			cell.SetString(RandStringBytes(6))
		}
	}
	if err := f.Save("tealeg.xlsx"); err != nil {
		fmt.Println(err)
		return
	}
	printBenchmarkInfo("tealeg/xlsx", startTime)
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
