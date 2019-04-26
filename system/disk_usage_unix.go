// +build !windows

package system

import (
	"os"

	"golang.org/x/sys/unix"
)

func (u *DefDiskUsage) Exists() (bool, error) {
	_, _, err := u.Stat()
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (u *DefDiskUsage) Stat() (uint64, uint64, error) {
	fd, err := os.Open(u.path)
	if err != nil {
		return 0, 0, err
	}
	var s unix.Statfs_t
	err = unix.Fstatfs(int(fd.Fd()), &s)
	return s.Blocks * uint64(s.Bsize), s.Bfree * uint64(s.Bsize), err
}
