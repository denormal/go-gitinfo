package gitinfo_test

import (
	"runtime"
	"testing"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestHere(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// does Here() return the same GitInfo as NewWithPath() with the path
	// set to the current file location?
	_, _file, _, _ok := runtime.Caller(0)
	if !_ok {
		t.Fatal("unexpected error; runtime.Caller() location not available")
	}

	_here, _err := gitinfo.Here()
	if _err != nil {
		t.Fatalf("unexpected error from Here(): %s", _err.Error())
	}
	_new, _err := gitinfo.NewWithPath(_file)
	if _err != nil {
		t.Fatalf("unexpected error from NewWithPath(): %s", _err.Error())
	}

	// we compare the two GitInfo structures by using their map representations
	//		- testing of Map will separately prove that this is sound
	_got := _here.Map()
	_expected := _new.Map()

	// ensure the maps are the same
	if len(_got) != len(_expected) {
		t.Fatalf(
			"unexpected map size; expected %d, got %d",
			len(_expected),
			len(_got),
		)
	}
	for _k, _v := range _expected {
		_value, _ok := _got[_k]
		if !_ok {
			t.Fatalf("missing value %q", _k)
		} else if _value != _v {
			t.Fatalf(
				"unexpected value for %q; expected %q, got %q",
				_k, _v, _value,
			)
		}
	}
} // TestHere()
