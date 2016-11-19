package minterm

import (
	"syscall"
	"unsafe"
)

type coord struct {
	x int16
	Y int16
}

type consoleScreenBufferInfo struct {
	size           coord
	cursorPosition coord
	attributes     uint16
	window         struct {
		left   int16
		top    int16
		right  int16
		bottom int16
	}
	maximumWindowSize coord
}

var getConsoleScreenBufferInfo = syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleScreenBufferInfo")

func TerminalSize() (columns, rows int, err error) {
	var csbi consoleScreenBufferInfo
	handle, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return
	}

	r1, _, lastErr := getConsoleScreenBufferInfo.Call(uintptr(handle), uintptr(unsafe.Pointer(&csbi)))
	if r1 == 0 {
		if lastErr == nil {
			err = syscall.EINVAL
		}
		err = lastErr
		return
	}

	columns = int(csbi.window.right - csbi.window.left + 1)
	rows = int(csbi.window.bottom - csbi.window.top + 1)
	return
}
