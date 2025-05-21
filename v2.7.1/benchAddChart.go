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

func benchAddChart(row, col int) {
	runtime.GC()
	startTime, f := time.Now(), excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	for r := 1; r <= row; r++ {
		for c := 0; c < col; c++ {
			if err := f.AddChart("Sheet1", "E1", &excelize.Chart{
				Type: excelize.Col3DClustered,
				Series: []excelize.ChartSeries{
					{
						Name:       "Sheet1!$A$2",
						Categories: "Sheet1!$B$1:$D$1",
						Values:     "Sheet1!$B$2:$D$2",
					},
					{
						Name:       "Sheet1!$A$3",
						Categories: "Sheet1!$B$1:$D$1",
						Values:     "Sheet1!$B$3:$D$3",
					},
					{
						Name:       "Sheet1!$A$4",
						Categories: "Sheet1!$B$1:$D$1",
						Values:     "Sheet1!$B$4:$D$4",
					}},
				Title: excelize.ChartTitle{
					Name: "Fruit 3D Clustered Column Chart",
				},
			}); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fileName := fmt.Sprintf("AddChart_r%dxc%d.xlsx", row, col)
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
	printBenchmarkInfo(fileName, startTime)
}
