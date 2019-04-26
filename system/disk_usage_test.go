package system

import (
	"testing"

	"github.com/aelsabbahy/goss/util"
)

func TestDiskUsageOK(t *testing.T) {
	u := NewDefDiskUsage("/", nil, util.Config{})
	total, free, err := u.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if total < free {
		t.Fatalf("total(%v) is less than free(%v)", total, free)
	}
	if total <= 0 {
		t.Fatalf("total(%v) <= 0", total)
	}

	exist, err := u.Exists()
	if !exist {
		t.Fatal("/ does not exist")
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestDiskUsageInvalid(t *testing.T) {
	u := NewDefDiskUsage("INVALID DIRECTORY", nil, util.Config{})
	_, _, err := u.Stat()
	if err == nil {
		t.Fatal("Stat should fail on invalid directory")
	}

	exist, err := u.Exists()
	if exist {
		t.Fatal("'INVALID DIRECTORY' existence check succeeded (and it should not have)")
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestUtilization(t *testing.T) {
	u := NewDefDiskUsage("DUMMY", nil, util.Config{})
	if u.UtilizationPercent(0, 0) != 100 {
		t.Fatal("DiskUsage.Utilization should report 100% if disk has no space")
	}
	ut := u.UtilizationPercent(100, 80)
	if ut != 20 {
		t.Fatalf("Utilization incorrect, got: %v, want: 20", ut)
	}
}
