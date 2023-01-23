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

func benchSetCellHyperLink(row, col int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	for r := 1; r <= row; r++ {
		for c := 1; c <= col; c++ {
			f.SetCellHyperLink("Sheet1", fmt.Sprintf("%s%d", excelize.ToAlphaString(c), r),
				"https://github.com/xuri/excelize", "External")
		}
	}
	fileName := fmt.Sprintf("SetCellHyperLink_r%dxc%d.xlsx", row, col)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
	printBenchmarkInfo(fileName, startTime)
}
