// Copyright 2021-2023 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license.
//
// This is a benchmark script for the Go language Spreadsheet (Excel / XLSX)
// library excelize: https://github.com/xuri/excelize

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	_ "image/jpeg"
)

var (
	funcFlag string
	rowsFlag int
	colsFlag int
	numFlag  int
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.StringVar(&funcFlag, "func", "", "function to benchmark")
	flag.IntVar(&rowsFlag, "rows", 0, "rows to benchmark")
	flag.IntVar(&colsFlag, "cols", 0, "columns to benchmark")
	flag.IntVar(&numFlag, "n", 0, "number to benchmark")
}

func main() {
	flag.Parse()
	if funcFlag == "" {
		fmt.Println("func is required flag")
		return
	}
	switch funcFlag {
	case "SetSheetRow":
		benchSetSheetRow(rowsFlag, colsFlag, numFlag)
	case "AddChart":
		benchAddChart(rowsFlag, colsFlag)
	case "SetCellHyperLink":
		benchSetCellHyperLink(rowsFlag, colsFlag)
	case "AddPicture":
		benchAddPicture(rowsFlag, colsFlag)
	case "RowIterator":
		benchRowIterator(rowsFlag, colsFlag)
	case "GetRows":
		benchGetRows(rowsFlag, colsFlag)
	case "MergeCell":
		benchMergeCell(numFlag)
	default:
		fmt.Println("unsupport benchmark function")
	}
}
