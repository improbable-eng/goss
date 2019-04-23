// +build !windows

package system

import (
	"github.com/aelsabbahy/goss/util"
)

type ServiceWindows struct{}

func NewServiceWindows(_ string, _ *System, _ util.Config) Service {
	panic("ServiceWindows used on non-windows platform")
}

func (_ *ServiceWindows) Service() string {
	panic("ServiceWindows used on non-windows platform")
}

func (s *ServiceWindows) Exists() (bool, error) {
	panic("ServiceWindows used on non-windows platform")
}

func (s *ServiceWindows) Enabled() (bool, error) {
	panic("ServiceWindows used on non-windows platform")
}

func (s *ServiceWindows) Running() (bool, error) {
	panic("ServiceWindows used on non-windows platform")
}
