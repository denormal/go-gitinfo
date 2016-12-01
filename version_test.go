package gitinfo_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

var _VERSION *regexp.Regexp

func TestVersion(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// otherwise, attempt to retrieve the version
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	}
	_version, _err := _info.Version()
	if _err != nil {
		t.Fatalf("unexpected Version() error: %s", _err.Error())
	} else if _version == "" {
		t.Fatalf("unexpected empty git version: %q", _version)
	} else if !_VERSION.Match([]byte(_version)) {
		t.Fatalf(
			"unexpected version; expected dotted-decimal, got %q",
			_version,
		)
	}

	// make sure the behaviour doesn't change with a path
	_cwd, _err := os.Getwd()
	if _err != nil {
		t.Fatalf(
			"unable to determine current working directory: %s",
			_err.Error(),
		)
	}
	for _, _path := range []string{"", _cwd} {
		_info, _err := gitinfo.New()
		if _err != nil {
			t.Fatalf("unexpected error from New(): %s", _err.Error())
		}
		_v, _err := _info.Version()
		if _err != nil {
			t.Fatalf("unexpected Version() error: %s", _err.Error())
		} else if _version == "" {
			t.Fatalf("unexpected empty git version: %q", _version)
		} else if !_VERSION.Match([]byte(_version)) {
			t.Fatalf(
				"unexpected version; expected dotted-decimal, got %q",
				_version,
			)
		} else if _v != _version {
			t.Fatalf(
				"%q: unexpected version; expected %q, got %q",
				_path, _version, _v,
			)
		}
	}
} // TestVersion()

func init() {
	_VERSION = regexp.MustCompile("^\\d+(\\.\\d+)*$")
} // init()
