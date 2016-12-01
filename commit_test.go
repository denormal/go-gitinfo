package gitinfo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestCommit(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// are we in a working copy?
	_working, _err := gittools.InWorkingCopy()
	if _err != nil {
		t.Fatalf("unexpected working copy error: %s", _err.Error())
	}

	// does GitInfo report the correct commit hash?
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	} else if _working {
		_commit, _err := _info.Commit()
		if _err != nil {
			t.Fatalf("unexpected error from Commit(): %s", _err.Error())
		} else if _commit == nil {
			t.Fatal("unexpected nil commit")
		} else if _commit.String() == "" {
			t.Fatal("unexpected empty commit")
		}
	} else {
		_commit, _err := _info.Commit()
		if _err == nil {
			t.Fatal("expected error from Commit(); none found")
		} else if _commit != nil {
			t.Fatalf("unexpected non-nil commit: %v", _commit)
		}
	}

	// ensure Prefix() behaves as expected
	if _working {
		_commit, _err := _info.Commit()
		if _err != nil {
			t.Fatalf("unexpected error from Commit(): %s", _err.Error())
		}
		_string := _commit.String()
		for _i := 0; _i <= len(_string); _i++ {
			_expected := _string[:_i]
			_prefix := _commit.Prefix(_i)
			if _prefix != _expected {
				t.Fatalf(
					"unexpected commit prefix; "+
						"expected %q, got %q for length %d",
					_expected, _prefix, _i,
				)
			}
		}
	}

	// ensure the commit is reported as nil if the GitInfo is
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
		_commit, _err := _info.Commit()
		if _err != nil {
			if _err != gitinfo.MissingWorkingCopyError {
				t.Fatalf("unexpected error from Commit(): %s", _err.Error())
			}
		}
		if _commit != nil {
			t.Fatalf(
				"unexpected commit; expected %v, got %v",
				nil, _commit,
			)
		}
	}
} // TestCommit()
