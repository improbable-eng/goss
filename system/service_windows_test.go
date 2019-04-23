package system

import (
	"testing"

	"github.com/aelsabbahy/goss/util"
	"golang.org/x/sys/windows/svc/mgr"
)

func cleanupTestSvc(t *testing.T, svcName string, m *mgr.Mgr) {
	svc, err := m.OpenService(svcName)
	if err != nil {
		// We assume that the service does not exist.
		return
	}
	if err = svc.Delete(); err != nil {
		t.Fatal(err)
	}
}

func TestDetectServiceWinMgr(t *testing.T) {
	m, err := mgr.Connect()
	if err != nil {
		t.Fatal(err)
	}

	// Create and auto-cleanup our service.
	const svcName = "randomGossTestService"
	cleanupTestSvc(t, svcName, m)
	mSvc, err := m.CreateService(svcName, "powershell", mgr.Config{})
	if err != nil {
		t.Fatal(err)
	}
	defer mSvc.Delete()

	svc := NewServiceWindows(svcName, nil, util.Config{})

	// Service should exist.
	ex, err := svc.Exists()
	if err != nil {
		t.Fatal(err)
	}
	if !ex {
		t.Fatalf("service %q not found", svcName)
	}

	// But it should not be running (and an attempt to run it would fail).
	running, err := svc.Running()
	if err != nil {
		t.Fatal(err)
	}
	if running {
		t.Fatalf("service %q is running", svcName)
	}
}

func TestDetectServiceWinMgrNoExist(t *testing.T) {
	svcName := "randomGossServiceThatShouldNotExist"
	svc := NewServiceWindows(svcName, nil, util.Config{})
	ex, err := svc.Exists()
	if err != nil {
		t.Fatal(err)
	}
	if ex {
		t.Fatalf("service %q should not exist", svcName)
	}
}
