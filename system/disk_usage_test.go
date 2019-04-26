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
}

func TestDiskUsageInvalid(t *testing.T) {
	u := NewDefDiskUsage("INVALID DIRECTORY", nil, util.Config{})
	_, _, err := u.Stat()
	if err == nil {
		t.Fatal("Stat should fail on invalid directory")
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
