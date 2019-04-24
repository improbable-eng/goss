package system

import (
	"syscall"

	"github.com/aelsabbahy/goss/util"
)

type DiskUsage interface {
	Exists() (bool, error)
	Path() string
	Utilization() (int, error)
}

type DefDiskUsage struct {
	path string
}

func NewDefDiskUsage(path string, system *System, config util.Config) DiskUsage {
	return &DefDiskUsage{
		path: path,
	}
}

// TODO(stefan): what is exists? good for?
func (u *DefDiskUsage) Exists() (bool, error) {
	return true, nil
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
