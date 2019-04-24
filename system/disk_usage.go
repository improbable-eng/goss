package system

import (
	"github.com/aelsabbahy/goss/util"
	"github.com/ricochet2200/go-disk-usage/du"
)

type DiskUsage interface {
	Path() string
	Utilization() string
}

type DefDiskUsage struct {
	path string
	dusg *du.DiskUsage
}

func NewDefDiskUsage(path string, system *System, config util.Config) DiskUsage {
	return &DefDiskUsage{
		path: path,
		dusg: du.NewDiskUsage(path),
	}
}

func (u *DefDiskUsage) Path() string {
	return u.path
}

func (u *DefDiskUsage) Utilization() string {
	return u.path
}
