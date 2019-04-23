package system

import (
	"fmt"
	"os"

	"github.com/aelsabbahy/goss/util"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

type ServiceWindows struct {
	service string
}

func NewServiceWindows(service string, system *System, config util.Config) Service {
	return &ServiceWindows{service: service}
}

func (s *ServiceWindows) Service() string {
	return s.service
}

func (s *ServiceWindows) Exists() (bool, error) {
	m, err := mgr.Connect()
	if err != nil {
		return false, err
	}
	svcs, err := m.ListServices()
	if err != nil {
		return false, err
	}
	for _, svc := range svcs {
		fmt.Fprintf(os.Stderr, "Found service %v\n", svc)
		if svc == s.service {
			return true, nil
		}
	}
	return false, nil
}

func (s *ServiceWindows) Enabled() (bool, error) {
	// Existence check.
	ex, err := s.Exists()
	if err != nil {
		return false, err
	}
	if !ex {
		return false, nil
	}

	m, err := mgr.Connect()
	if err != nil {
		return false, err
	}
	svc, err := m.OpenService(s.service)
	if err != nil {
		return false, err
	}
	cfg, err := svc.Config()
	if err != nil {
		return false, err
	}
	return cfg.StartType == mgr.StartAutomatic, nil
}

func (s *ServiceWindows) Running() (bool, error) {
	// Existence check.
	ex, err := s.Exists()
	if err != nil {
		return false, err
	}
	if !ex {
		return false, nil
	}

	m, err := mgr.Connect()
	if err != nil {
		return false, err
	}
	svc, err := m.OpenService(s.service)
	if err != nil {
		return false, err
	}
	q, err := svc.Query()
	if err != nil {
		return false, err
	}
	fmt.Fprintf(os.Stderr, "Running=%+v\n", q)
	return q.State == windows.SERVICE_RUNNING, nil
}
