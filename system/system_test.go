package system

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/aelsabbahy/goss/util"
)

type noInputs func() string

// test that a function with no inputs returns one of the expected strings
func testOutputs(f noInputs, validOutputs []string, t *testing.T) {
	output := f()
	// use reflect to get the name of the function
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	failed := true
	for _, valid := range validOutputs {
		if output == valid {
			failed = false
		}
	}
	if failed {
		t.Errorf("Function %v returned %v, which is not one of %v", name, output, validOutputs)
	}
}

func TestPackageManager(t *testing.T) {
	t.Parallel()
	testOutputs(
		DetectPackageManager,
		[]string{"deb", "rpm", "apk", "pacman", ""},
		t,
	)
}

func TestDetectService(t *testing.T) {
	t.Parallel()
	testOutputs(
		DetectService,
		[]string{"systemd", "init", "alpineinit", "upstart", "windows", ""},
		t,
	)
}

func TestDetectServiceWin(t *testing.T) {
	if runtime.GOOS != "windows" {
		return
	}
	t.Parallel()
	svcName := "hello123_this_should_not_exist"
	svc := NewServiceWindows(svcName, nil, util.Config{})
	ex, err := svc.Exists()
	if err != nil {
		t.Fatal(err)
	}
	if ex {
		t.Fatalf("service %q should not exist", svcName)
	}
}

func TestDetectServiceWin2(t *testing.T) {
	t.Fatalf("failed2")
}

func TestDetectDistro(t *testing.T) {
	t.Parallel()
	testOutputs(
		DetectDistro,
		[]string{"ubuntu", "redhat", "alpine", "arch", "debian", ""},
		t,
	)
}

func TestHasCommand(t *testing.T) {
	t.Parallel()
	if !HasCommand("sh") {
		t.Error("System didn't have sh!")
	}
}
