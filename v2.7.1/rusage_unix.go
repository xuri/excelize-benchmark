// Copyright 2021-2026 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license.
//
// This is a benchmark script for the Go language Spreadsheet (Excel / XLSX)
// library excelize: https://github.com/xuri/excelize

//go:build !windows

package main

import "syscall"

// maxRSS returns the maximum resident set size of the current process.
func maxRSS() uint64 {
	var rusage syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &rusage)
	return uint64(rusage.Maxrss)
}
