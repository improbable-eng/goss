// +build !windows

package system

import (
	"os"

	"golang.org/x/sys/unix"
)

func (u *DefDiskUsage) Exists() (bool, error) {
	_, err := u.stat()
	if err == nil {
		return true, nil
	}
	if errS, ok := err.(unix.Errno); ok && errS == unix.ENOENT {
		return false, nil
	}
	return false, err
}

func (u *DefDiskUsage) Path() string {
	return u.path
}

func (u *DefDiskUsage) TotalBytes() (uint64, error) {
	s, err := u.stat()
	if err != nil {
		return 0, err
	}
	return s.Blocks * uint64(s.Bsize), nil
}

func (u *DefDiskUsage) FreeBytes() (uint64, error) {
	s, err := u.stat()
	if err != nil {
		return 0, err
	}
	return s.Bfree * uint64(s.Bsize), nil
}

func (u *DefDiskUsage) Utilization() (int, error) {
	s, err := u.stat()
	if err != nil {
		return 0, err
	}
	return int(100 * (1 - float32(s.Bfree)/float32(s.Blocks))), nil
}

func (u *DefDiskUsage) stat() (*unix.Statfs_t, error) {
	fd, err := os.Open(u.path)
	if err != nil {
		return nil, err
	}
	var s unix.Statfs_t
	err = unix.Fstatfs(int(fd.Fd()), &s)
	return &s, err
}
