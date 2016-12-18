package gitinfo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestBranch(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// are we in a working copy?
	_working, _err := gittools.InWorkingCopy()
	if _err != nil {
		t.Fatalf("unexpected working copy error: %s", _err.Error())
	}

	// does GitInfo report the correct branch?
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	} else if _working {
		_branch, _err := _info.Branch()
		if _err != nil {
			t.Fatalf("unexpected error from Branch(): %s", _err.Error())
		} else if _branch == "" {
			t.Fatal("unexpected empty branch")
		}
	} else {
		_branch, _err := _info.Branch()
		if _err == nil {
			t.Fatal("expected error from Branch(); none found")
		} else if _branch != "" {
			t.Fatalf("unexpected non-nil commit: %q", _branch)
		}
	}

	// ensure the branch is reported as empty if the GitInfo is
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
	} else {
		_branch, _err := _info.Branch()
		if _err != nil {
			if _err != gitinfo.MissingWorkingCopyError {
				t.Fatalf("unexpected error from Branch(): %s", _err.Error())
			}
		}
		if _branch != "" {
			t.Fatalf(
				"unexpected commit; expected %q, got %q",
				"", _branch,
			)
		}
	}
} // TestBranch()
