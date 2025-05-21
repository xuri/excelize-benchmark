// Copyright 2021-2025 The excelize Authors. All rights reserved. Use of
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

func benchSetSheetRow(row, col, cellLen int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	for r := 1; r <= row; r++ {
		row := make([]interface{}, col)
		for c := 0; c < col; c++ {
			row[c] = randStringBytes(cellLen)
		}
		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", r), &row)
	}
	fileName := fmt.Sprintf("SetSheetRow_r%dxc%d.xlsx", row, col)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return
	}
	printBenchmarkInfo(fileName, startTime)
}
