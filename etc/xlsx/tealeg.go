package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"syscall"
	"time"

	"github.com/tealeg/xlsx"
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
	for i := 0; i < 102400; i++ {
		row := sheet.AddRow()
		for i := 0; i < 50; i++ {
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
	var rusage syscall.Rusage
	var bToMb = func(b uint64) uint64 {
		return b / 1024 / 1024
	}
	runtime.ReadMemStats(&memStats)
	syscall.Getrusage(syscall.RUSAGE_SELF, &rusage)
	fmt.Printf("Func: %s \tRSS = %v MB\tAlloc = %v MB\tTotalAlloc = %v MB\tSys = %v MB\tNumGC = %v \tCost = %s\n",
		fn, bToMb(uint64(rusage.Maxrss)), bToMb(memStats.Alloc), bToMb(memStats.TotalAlloc), bToMb(memStats.Sys), memStats.NumGC, time.Since(startTime))
}
