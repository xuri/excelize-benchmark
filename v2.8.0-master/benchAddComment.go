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

	"github.com/xuri/excelize/v2"
)

func benchAddComment(row, col, cellLen int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for r := 1; r <= row; r++ {
		for c := 1; c <= col; c++ {
			cell, err := excelize.CoordinatesToCellName(c, r)
			if err != nil {
				fmt.Println(err)
				return
			}
			if err = f.AddComment("Sheet1", excelize.Comment{
				Cell: cell,
				Text: randStringBytes(cellLen),
			}); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fileName := fmt.Sprintf("AddComment_r%dxc%d.xlsx", row, col)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return
	}
	printBenchmarkInfo(fileName, startTime)
}
