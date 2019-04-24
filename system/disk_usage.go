package system

import "github.com/aelsabbahy/goss/util"

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
