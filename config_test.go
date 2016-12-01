package gitinfo_test

import (
	"testing"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestConfig(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// otherwise, attempt to retrieve the configuration
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	}

	// ensure we have a non-nil configuration
	_config := _info.Config()
	if _config == nil {
		t.Fatal("unexpected nil git configuration")
	}
	_all := _config.All()
	if len(_all) == 0 {
		t.Fatal("unexpected empty git configuration")
	}
} // TestConfig()
