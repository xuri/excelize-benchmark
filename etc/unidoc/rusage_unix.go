//go:build !windows

package main

import "syscall"

// maxRSS returns the maximum resident set size of the current process.
func maxRSS() uint64 {
	var rusage syscall.Rusage
	syscall.Getrusage(syscall.RUSAGE_SELF, &rusage)
	return uint64(rusage.Maxrss)
}
