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

func benchGetRows(row, col int) {
	runtime.GC()
	startTime := time.Now()
	f, err := excelize.OpenFile(fmt.Sprintf("StreamWriter_r%dxc%d.xlsx", row, col))
	if err != nil {
		fmt.Println(err)
		return
	}
	c := 0
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, r := range rows {
		for _, colCell := range r {
			_ = colCell
			c++
		}
	}
	if c != row*col {
		fmt.Println("Test GetRows Error", c, row*col)
		return
	}
	printBenchmarkInfo(fmt.Sprintf("GetRows_r%dxc%d.xlsx", row, col), startTime)
}
