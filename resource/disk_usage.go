package resource

import (
	"github.com/aelsabbahy/goss/system"
	"github.com/aelsabbahy/goss/util"
)

type DiskUsage struct {
	Title       string  `json:"title,omitempty" yaml:"title,omitempty"`
	Meta        meta    `json:"meta,omitempty" yaml:"meta,omitempty"`
	Path        string  `json:"-" yaml:"-"`
	TotalBytes  matcher `json:"total_bytes" yaml:"utilization"`
	FreeBytes   matcher `json:"free_bytes" yaml:"utilization"`
	Utilization matcher `json:"utilization" yaml:"utilization"`
}

func (u *DiskUsage) ID() string      { return u.Path }
func (u *DiskUsage) SetID(id string) { u.Path = id }

func (u *DiskUsage) GetTitle() string { return u.Title }
func (u *DiskUsage) GetMeta() meta    { return u.Meta }

func (u *DiskUsage) Validate(sys *system.System) []TestResult {
	skip := false
	sysDiskUsage := sys.NewDiskUsage(u.Path, sys, util.Config{})

	var results []TestResult
	results = append(results, ValidateValue(u, "utilization", u.Utilization, sysDiskUsage.Utilization, skip))
	return results
}

func NewDiskUsage(sysDiskUsage system.DiskUsage, config util.Config) (*DiskUsage, error) {
	path := sysDiskUsage.Path()
	utilization, err := sysDiskUsage.Utilization()
	return &DiskUsage{
		Path:        path,
		Utilization: utilization,
	}, err
}
