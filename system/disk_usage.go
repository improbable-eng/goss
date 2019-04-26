package system

import "github.com/aelsabbahy/goss/util"

type DiskUsage interface {
	Exists() (bool, error)
	Path() string
	// Utilization returns utilization from percent. totalBytes and freeBytes come from Stat().
	UtilizationPercent(totalBytes uint64, freeBytes uint64) int
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

func (u *DefDiskUsage) UtilizationPercent(totalBytes, freeBytes uint64) int {
	if totalBytes == 0 {
		// If totalBytes is 0, set utilization to 100%. Not sure if this can ever happen, but it protects us from division by zero.
		return 100
	}
	return 100 - int(freeBytes*100/totalBytes)
}
