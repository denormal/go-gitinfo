package gitinfo_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestEditor(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// run the editor test
	editor(t)

	// reset the environment and rerun the tests
	defer env("GIT_EDITOR", "")
	defer env("EDITOR", "")
	defer env("VISUAL", "")
	t.Run("EmptyEnv", func(t *testing.T) { editor(t) })

	// ensure the environment variables are correctly honoured
	_value := func() string {
		return strconv.FormatInt(time.Now().UnixNano(), 16)
	}

	//		- GIT_EDITOR
	_target := _value()
	env("GIT_EDITOR", _target)
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	} else if _info.Editor() != _target {
		t.Fatalf(
			"unexpected editor; expected %q, got %q",
			_target, _info.Editor(),
		)
	}
	t.Run("GIT_EDITOR", func(t *testing.T) { editor(t) })

	//		- EDITOR
	env("GIT_EDITOR", "")
	_target = _value()
	env("EDITOR", _target)
	_info, _err = gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	} else if _info.Editor() != _target {
		t.Fatalf(
			"unexpected editor; expected %q, got %q",
			_target, _info.Editor(),
		)
	}
	t.Run("EDITOR", func(t *testing.T) { editor(t) })

	//		- VISUAL
	env("EDITOR", "")
	_target = _value()
	env("VISUAL", _target)
	_info, _err = gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	} else if _info.Editor() != _target {
		t.Fatalf(
			"unexpected editor; expected %q, got %q",
			_target, _info.Editor(),
		)
	}
	t.Run("VISUAL", func(t *testing.T) { editor(t) })
} // TestEditor()

//
// helper methods
//

func editor(t *testing.T) {

	// otherwise, attempt to retrieve the configuration
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	}

	// extract the editor value
	_editor := _info.Editor()

	// ensure this value is not empty
	if _editor == "" {
		t.Fatalf(
			"unexpected editor; expected non-empty value, got %q", _editor,
		)
	}

	// examine the environment for the preferred editor
	//		- these are defined in precedence order
	//		- if we get a match then we know it should be the expected
	//		  editor value
	for _, _env := range []string{"GIT_EDITOR", "EDITOR", "VISUAL"} {
		_expected := os.Getenv(_env)
		if _expected != "" {
			if _expected != _editor {
				t.Fatalf(
					"unexpected editor; expected %q, got %q",
					_expected, _editor,
				)
			}
			// we matched the editor, so we are done with our test
			return
		}
	}

	// the editor was not defined in the environment, so examine the config
	_config := _info.Config()
	if _config == nil {
		t.Fatal("unexpected nil git configuration")
	}
	_expected := _config.Get("core.editor")
	if _expected != nil {
		if _expected.String() != "" {
			if _expected.String() != _editor {
				t.Fatalf(
					"unexpected editor; expected %q, got %q",
					_expected, _editor,
				)
			}
		}
	}

	// there's no editor defined in the configuration, so we should have
	// the default value
	if _editor != "vi" {
		t.Fatalf(
			"unexpected editor; expected %q, got %q",
			_expected, _editor,
		)
	}
} // editor()

func env(name, value string) (func(), error) {
	_previous := os.Getenv(name)
	_callback := func() {
		if _previous == "" {
			os.Unsetenv(name)
		} else {
			os.Setenv(name, _previous)
		}
	}

	// attempt to set the new environment value
	return _callback, os.Setenv(name, value)
} // env()
