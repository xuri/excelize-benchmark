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

func benchAddPicture(row, col int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	for r := 1; r <= row; r++ {
		for c := 1; c <= col; c++ {
			cell, err := excelize.CoordinatesToCellName(c, r)
			if err != nil {
				fmt.Println(err)
				return
			}
			// Insert a picture.
			if err := f.AddPicture("Sheet1", cell, "excel.jpg", ""); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fileName := fmt.Sprintf("AddPicture_r%dxc%d.xlsx", row, col)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
	printBenchmarkInfo(fileName, startTime)
}
