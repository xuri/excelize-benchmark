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

func benchSetCellHyperLink(row, col int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	for r := 1; r <= row; r++ {
		for c := 1; c <= col; c++ {
			cell, err := excelize.CoordinatesToCellName(c, r)
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := f.SetCellHyperLink("Sheet1", cell,
				"https://github.com/xuri/excelize", "External"); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fileName := fmt.Sprintf("SetCellHyperLink_r%dxc%d.xlsx", row, col)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
	printBenchmarkInfo(fileName, startTime)
}
