package gitinfo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestRoot(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// are we in a working copy?
	//		- it's likely we are since getting this code will usually
	//	      be the result of a git clone
	_working, _err := gittools.WorkingCopy("")
	if _err != nil {
		if _err == gittools.MissingWorkingCopyError {
			// ensure we report no working copy
			_info, _err := gitinfo.New()
			if _err != nil {
				t.Fatalf(
					"unexpected error from gitinfo.New(): %s",
					_err.Error(),
				)
			} else if _info.Root() != "" {
				t.Fatalf(
					"working root mismatch; expected %q, got %q",
					"", _info.Root(),
				)
			}
		}
	}

	// ensure the working copy is as expected
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("New() %q", _err.Error())
	} else if _info.Root() != _working {
		t.Fatalf(
			"working root mismatch; expected %q, got %q",
			_working, _info.Root(),
		)
	}

	// ensure the working directory is reported as "" if the GitInfo is
	// created in a folder that is not a working copy

	// create a temporary directory
	//		- this should not be a working copy
	_dir, _err := ioutil.TempDir("", "")
	if _err != nil {
		t.Fatalf("unable to create temporary directory: %s", _err.Error())
	}
	defer os.RemoveAll(_dir)

	_info, _err = gitinfo.NewWithPath(_dir)
	if _err != nil {
		t.Fatalf("%q: unexpected error from New(): %s", _dir, _err.Error())
	} else if _info.Root() != "" {
		t.Fatalf(
			"unexpected root; expected %q, got %q",
			"", _info.Root(),
		)
	}
} // TestRoot()
