package gitinfo_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/denormal/go-gitinfo"
	"github.com/denormal/go-gittools"
)

func TestUser(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// reset the environment
	defer unset(t, "GIT_AUTHOR_NAME")
	defer unset(t, "GIT_COMMITTER_NAME")
	defer unset(t, "GIT_AUTHOR_EMAIL")
	defer unset(t, "GIT_COMMITTER_EMAIL")
	defer unset(t, "EMAIL")

	// otherwise, attempt to retrieve the configuration
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	}

	// extract the user form the info instance
	_user := _info.User()
	if _user == nil {
		t.Fatalf("unexpected nil user")
	}

	// ensure we have a non-nil configuration
	_config := _info.Config()
	if _config == nil {
		t.Fatal("unexpected nil git configuration")
	}

	// ensure the user name is as expected
	_name := _user.Name()
	_expected := _config.Get("user.name")
	if _expected == nil {
		if _name != "" {
			t.Errorf(
				"unexpected user name; expected %q, got %q",
				"", _name,
			)
		}
	} else if _name != _expected.String() {
		t.Errorf(
			"unexpected user name; expected %q, got %q",
			_expected.String(), _name,
		)
	}

	// ensure the user email is as expected
	_email := _user.Email()
	_expected = _config.Get("user.email")
	if _expected == nil {
		if _email != "" {
			t.Errorf(
				"unexpected user email; expected %q, got %q",
				"", _email,
			)
		}
	} else if _email != _expected.String() {
		t.Errorf(
			"unexpected user email; expected %q, got %q",
			_expected.String(), _email,
		)
	}

	// ensure the string version of the user is as expected
	if _name == "" {
		if _user.String() != _email {
			t.Errorf(
				"unexpected user string; expected %q, got %q",
				_email, _user.String(),
			)
		}
	} else if _email == "" {
		if _user.String() != _name {
			t.Errorf(
				"unexpected user string; expected %q, got %q",
				_name, _user.String(),
			)
		}
	} else {
		_string := fmt.Sprintf("%s <%s>", _name, _email)
		if _user.String() != _string {
			t.Errorf(
				"unexpected user string; expected %q, got %q",
				_string, _user.String(),
			)
		}
	}
} // TestUser()

func TestUserEnv(t *testing.T) {
	// if we don't have git installed, then skip this test
	if !gittools.HasGit() {
		t.Skip("git not installed")
	}

	// clear the environment and forcibly set the user's name and e-mail
	// through environment variables
	defer unset(t, "GIT_AUTHOR_NAME")
	defer unset(t, "GIT_COMMITTER_NAME")
	defer unset(t, "GIT_AUTHOR_EMAIL")
	defer unset(t, "GIT_COMMITTER_EMAIL")

	// ensure GIT_COMMITTER_NAME & GIT_COMMITTER_EMAIL are honoured
	_name := strconv.FormatInt(time.Now().UnixNano(), 16)
	_email := strconv.FormatInt(time.Now().UnixNano(), 16)
	set(t, "GIT_COMMITTER_NAME", _name)
	set(t, "GIT_COMMITTER_EMAIL", _email)

	// otherwise, attempt to retrieve the configuration
	_info, _err := gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	}

	// extract the user form the info instance
	_user := _info.User()
	if _user == nil {
		t.Fatalf("unexpected nil user")
	} else if _user.Name() != _name {
		t.Fatalf(
			"GIT_COMMITTER_NAME: unexpected user name; expected %q, got %q",
			_name, _user.Name(),
		)
	} else if _user.Email() != _email {
		t.Fatalf(
			"GIT_COMMITTER_EMAIL: unexpected user email; expected %q, got %q",
			_email, _user.Email(),
		)
	}

	// ensure GIT_AUTHOR_NAME & GIT_AUTHOR_EMAIL are honoured
	//		- these should take precedence over the COMMITTER values
	_name = strconv.FormatInt(time.Now().UnixNano(), 16)
	_email = strconv.FormatInt(time.Now().UnixNano(), 16)
	set(t, "GIT_AUTHOR_NAME", _name)
	set(t, "GIT_AUTHOR_EMAIL", _email)

	// otherwise, attempt to retrieve the configuration
	_info, _err = gitinfo.New()
	if _err != nil {
		t.Fatalf("unexpected error from New(): %s", _err.Error())
	}

	// extract the user form the info instance
	_user = _info.User()
	if _user == nil {
		t.Fatalf("unexpected nil user")
	} else if _user.Name() != _name {
		t.Fatalf(
			"GIT_AUTHOR_NAME: unexpected user name; expected %q, got %q",
			_name, _user.Name(),
		)
	} else if _user.Email() != _email {
		t.Fatalf(
			"GIT_AUTHOR_EMAIL: unexpected user email; expected %q, got %q",
			_email, _user.Email(),
		)
	}
} // TestUserEnv()

//
// helper function
//

func unset(t *testing.T, env string) func() {
	// retrieve the previous value
	_env := os.Getenv(env)
	_reset := func() {
		if _env != "" {
			os.Setenv(env, _env)
		}
	}

	// attempt to unset the variable
	_err := os.Unsetenv(env)
	if _err != nil {
		t.Fatalf("%s: unset failed: %s", env, _err.Error())
	}

	// return the reset callback
	return _reset
} // unset()

func set(t *testing.T, env, value string) func() {
	// retrieve the previous value
	_env := os.Getenv(env)
	_reset := func() {
		if _env != "" {
			os.Setenv(env, _env)
		}
	}

	// attempt to set the environment variable
	_err := os.Setenv(env, value)
	if _err != nil {
		t.Fatalf("%s: set failed: %s", env, _err.Error())
	}

	// return the reset callback
	return _reset
} // set()
