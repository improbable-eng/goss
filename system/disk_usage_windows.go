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

func (u *DefDiskUsage) Path() string {
	return u.path
}

func (u *DefDiskUsage) Utilization() (int, error) {
	totalBytes, availBytes, err := u.stat()
	if err != nil {
		return 0, err
	}
	return int(100 * (1 - float32(freeBytes)/float32(totalBytes))), nil
}

func (u *DefDiskUsage) stat() (uint64, uint64, error) {
	var dummy, totalBytes, availBytes uint64

	r1, _, err := windows.
		NewLazySystemDLL("kernel32.dll").
		NewProc("GetDiskFreeSpaceExW").
		Call(
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(u.path))),
			uintptr(unsafe.Pointer(&dummy)), // free bytes available to caller
			uintptr(unsafe.Pointer(&totalBytes)),
			uintptr(unsafe.Pointer(&availBytes)))

	if r1 == 0 {
		// syscall errors out if r1 is zero. err is always not nil.
		return 0, 0, fmt.Errorf("failed to call kernel32.dll:GetDiskFreeSpaceExW: %v", err)
	}

	return totalBytes, availBytes, nil
}
