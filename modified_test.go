package gitinfo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestModified(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// the current path should be a working copy since this is most
	// likely a clone of this module
	//		- if it's not a working copy, then skip this test
	_working, _err := gittools.WorkingCopy("")
	if _err != nil {
		t.Fatalf("unexpected working copy error: %s", _err.Error())
	} else if _working == "" {
		t.Skip("not in working copy")
	}

	// is this working copy reported as modified?
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	} else {
		_modified, _err := _info.Modified()
		if _err != nil {
			t.Fatalf("unexpected error from Modified(): %s", _err.Error())
		} else if !_modified {
			// since we report this working copy as modified, there's no
			// more testing we will do
			//		- in the future we may wish to be more thorough
			//		- that will likely involve creating a fake git repository
			//		  to test with
			//		- for now we keep things simple
			return
		}
	}

	// the working copy is currently unmodified
	//		- create a temporary file within the working copy and
	//		  ensure this changes the modified state
	_file, _err := ioutil.TempFile(_working, "")
	if _err != nil {
		t.Fatalf("unexpected error making temporary file: %s", _err.Error())
	}
	defer os.Remove(_file.Name())

	_modified, _err := _info.Modified()
	if _err != nil {
		t.Fatalf("unexpected error from Modified(): %s", _err.Error())
	} else if !_modified {
		t.Fatalf("modified mismatch; expected %v, got %v", true, _modified)
	}
} // TestModified()
