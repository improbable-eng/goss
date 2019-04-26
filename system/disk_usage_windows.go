package system

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func (u *DefDiskUsage) Exists() (bool, error) {
	_, _, _, err := u.stat()
	return err == nil, err
}

func (u *DefDiskUsage) Stat() (uint64, uint64, error) {
	var dummy, totalBytes, freeBytes uint64

	r1, _, err := windows.
		NewLazySystemDLL("kernel32.dll").
		NewProc("GetDiskFreeSpaceExW").
		Call(
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(u.path))),
			uintptr(unsafe.Pointer(&dummy)), // free bytes available to caller
			uintptr(unsafe.Pointer(&totalBytes)),
			uintptr(unsafe.Pointer(&freeBytes)))

	if r1 == 0 {
		// syscall errors out if r1 is zero. err is always not nil.
		return 0, 0, fmt.Errorf("failed to call kernel32.dll:GetDiskFreeSpaceExW: %v", err)
	}

	return totalBytes, freeBytes, nil
}
