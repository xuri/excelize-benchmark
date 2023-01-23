// Copyright 2021-2023 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license.
//
// This is a benchmark script for the Go language Spreadsheet (Excel / XLSX)
// library excelize: https://github.com/xuri/excelize

package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/xuri/excelize"
)

func benchMergeCell(n int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	for r := 1; r <= n; r++ {
		f.MergeCell("Sheet1", fmt.Sprintf("A%d", r), fmt.Sprintf("B%d", r))
	}
	fileName := fmt.Sprintf("MergeCell_r%d.xlsx", n)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
	printBenchmarkInfo(fileName, startTime)
}
