package system

import (
	"testing"

	"github.com/aelsabbahy/goss/util"
	"golang.org/x/sys/windows/svc/mgr"
)

const svcName = "randomGossTestService"

func cleanupTestSvc(t *testing.T, m *mgr.Mgr) {
	svc, err := m.OpenService(svcName)
	if err != nil {
		// We assume that the service does not exist.
		return
	}
	if err = svc.Delete(); err != nil {
		t.Fatal(err)
	}
}

func TestDetectServiceWinCR(t *testing.T) {
	m, err := mgr.Connect()
	if err != nil {
		t.Fatal(err)
	}

	// Create and auto-cleanup our service.
	cleanupTestSvc(t, m)
	mSvc, err := m.CreateService(svcName, "powershell", mgr.Config{})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = mSvc.Delete(); err != nil {
			t.Fatal(err)
		}
	}()

	svc := NewServiceWindows(svcName, nil, util.Config{})

	// Does the service exist?
	ex, err := svc.Exists()
	if err != nil {
		t.Fatal(err)
	}
	if !ex {
		t.Fatalf("service %q not found", svcName)
	}

	// Is it running?
	running, err := svc.Running()
	if err != nil {
		t.Fatal(err)
	}
	if !running {
		t.Fatalf("service %q is not running", svcName)
	}
}
