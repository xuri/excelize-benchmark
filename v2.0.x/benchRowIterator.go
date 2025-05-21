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

func benchRowIterator(row, col int) {
	runtime.GC()
	startTime := time.Now()
	f, err := excelize.OpenFile(fmt.Sprintf("SetSheetRow_r%dxc%d.xlsx", row, col))
	if err != nil {
		fmt.Println(err)
		return
	}
	c := 0
	rows, err := f.Rows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		for _, colCell := range row {
			_ = colCell
			c++
		}
	}
	if c != row*col {
		fmt.Println("Test Iterator Error")
		return
	}
	printBenchmarkInfo(fmt.Sprintf("RowIterator_r%dxc%d.xlsx", row, col), startTime)
}
