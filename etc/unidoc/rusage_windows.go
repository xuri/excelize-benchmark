//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

// processMemoryCounters maps the Windows PROCESS_MEMORY_COUNTERS structure.
type processMemoryCounters struct {
	cb                         uint32
	pageFaultCount             uint32
	peakWorkingSetSize         uintptr
	workingSetSize             uintptr
	quotaPeakPagedPoolUsage    uintptr
	quotaPagedPoolUsage        uintptr
	quotaPeakNonPagedPoolUsage uintptr
	quotaNonPagedPoolUsage     uintptr
	pagefileUsage              uintptr
	peakPagefileUsage          uintptr
}

var (
	psapi                    = syscall.NewLazyDLL("psapi.dll")
	procGetProcessMemoryInfo = psapi.NewProc("GetProcessMemoryInfo")
)

// maxRSS returns the peak working set size of the current process in bytes.
func maxRSS() uint64 {
	var counters processMemoryCounters
	counters.cb = uint32(unsafe.Sizeof(counters))
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return 0
	}
	ret, _, _ := procGetProcessMemoryInfo.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&counters)),
		uintptr(counters.cb),
	)
	if ret == 0 {
		return 0
	}
	return uint64(counters.peakWorkingSetSize)
}
