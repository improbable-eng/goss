package system

import "github.com/aelsabbahy/goss/util"

type DiskUsage interface {
	Exists() (bool, error)
	Path() string
	// Stat returns total bytes, free bytes and error (or nil).
	Stat() (uint64, uint64, error)
}

type DefDiskUsage struct {
	path string
}

func NewDefDiskUsage(path string, system *System, config util.Config) DiskUsage {
	return &DefDiskUsage{
		path: path,
	}
}

func (u *DefDiskUsage) Path() string {
	return u.path
}
