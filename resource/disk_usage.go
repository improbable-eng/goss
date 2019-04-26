package resource

import (
	"github.com/aelsabbahy/goss/system"
	"github.com/aelsabbahy/goss/util"
)

type DiskUsage struct {
	Title              string  `json:"title,omitempty" yaml:"title,omitempty"`
	Meta               meta    `json:"meta,omitempty" yaml:"meta,omitempty"`
	Path               string  `json:"-" yaml:"-"`
	Exists             matcher `json:"exists" yaml:"exists"`
	TotalBytes         matcher `json:"total_bytes" yaml:"total_bytes"`
	FreeBytes          matcher `json:"free_bytes" yaml:"free_bytes"`
	UtilizationPercent matcher `json:"utilization_percent" yaml:"utilization_percent"`
}

func (u *DiskUsage) ID() string      { return u.Path }
func (u *DiskUsage) SetID(id string) { u.Path = id }

func (u *DiskUsage) GetTitle() string { return u.Title }
func (u *DiskUsage) GetMeta() meta    { return u.Meta }

func (u *DiskUsage) Validate(sys *system.System) []TestResult {
	skip := false
	// TODO handle err?
	du, _ := NewDiskUsage(sys.NewDiskUsage(u.Path, sys, util.Config{}), util.Config{})

	results := []TestResult{ValidateValue(u, "exists", u.Exists, du.Exists, skip)}
	if shouldSkip(results) {
		skip = true
	}
	if u.TotalBytes != nil {
		results = append(results, ValidateValue(u, "total_bytes", u.TotalBytes, du.TotalBytes, skip))
	}
	if u.FreeBytes != nil {
		results = append(results, ValidateValue(u, "free_bytes", u.FreeBytes, du.FreeBytes, skip))
	}
	if u.UtilizationPercent != nil {
		results = append(results, ValidateValue(u, "utilization_percent", u.UtilizationPercent, du.UtilizationPercent, skip))
	}
	return results
}

func NewDiskUsage(sysDiskUsage system.DiskUsage, config util.Config) (*DiskUsage, error) {
	totalBytes, freeBytes, err := sysDiskUsage.Stat()
	if err != nil {
		return &DiskUsage{
			Path:   sysDiskUsage.Path(),
			Exists: false,
		}, nil
	}
	return &DiskUsage{
		Exists:             true,
		TotalBytes:         totalBytes,
		FreeBytes:          freeBytes,
		UtilizationPercent: sysDiskUsage.UtilizationPercent(totalBytes, freeBytes),
	}, nil
}
