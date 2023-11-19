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

func benchSetSheetRow(row, col, cellLen int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for r := 1; r <= row; r++ {
		row := make([]interface{}, col)
		for c := 0; c < col; c++ {
			row[c] = randStringBytes(cellLen)
		}
		cell, err := excelize.CoordinatesToCellName(1, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err = f.SetSheetRow("Sheet1", cell, &row); err != nil {
			fmt.Println(err)
			return
		}
	}
	fileName := fmt.Sprintf("SetSheetRow_r%dxc%d.xlsx", row, col)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return
	}
	printBenchmarkInfo(fileName, startTime)
}
