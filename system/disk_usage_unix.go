// +build !windows

package system

import (
	"syscall"
)

func (u *DefDiskUsage) Exists() (bool, error) {
	_, err := u.stat()
	if err == nil {
		return true, nil
	}
	if errS, ok := err.(syscall.Errno); ok && errS == syscall.ENOENT {
		return false, nil
	}
	return false, err
}

func (u *DefDiskUsage) Path() string {
	return u.path
}

func (u *DefDiskUsage) Utilization() (int, error) {
	s, err := u.stat()
	return int(100 * (1 - float32(s.Bfree)/float32(s.Blocks))), err
}

func (u *DefDiskUsage) stat() (syscall.Statfs_t, error) {
	var s syscall.Statfs_t
	err := syscall.Statfs(u.path, &s)
	return s, err
}
