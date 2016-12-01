package gitinfo

import (
	"fmt"
	"os"

	"github.com/denormal/go-gitconfig"
)

// User is the interface representing the git user.
type User interface {
	// Name returns the name of the current git user, or the empty string
	// if no name is configured.
	Name() string

	// Email returns the e-mail addrdd of the current git user, or the empty
	// string if no e-mail address is configured.
	Email() string

	// String() returns a string representation of the git user's name and
	// e-mail address, or the empty string if neither are defined.
	String() string
}

// user is the implementation of the User interface
type user struct {
	name  string
	email string
}

// newUser returns the user instance as detailed in the given git configuration
func newUser(config gitconfig.GitConfig) User {
	// extract the name & email from the git configuration
	_name := ""
	if config != nil {
		_config := config.Get("user.name")
		if _config != nil {
			_name = _config.String()
		}
	}
	_email := ""
	if config != nil {
		_config := config.Get("user.email")
		if _config != nil {
			_email = _config.String()
		}
	}

	return &user{name: _name, email: _email}
} // newUser()

// Name returns the name of the current git user, or the empty string
// if no name is configured.
func (u *user) Name() string {
	// is the git user name defined in the environment?
	_n := os.Getenv("GIT_AUTHOR_NAME")
	if _n == "" {
		_n = os.Getenv("GIT_COMMITTER_NAME")
		if _n == "" {
			_n = u.name
		}
	}

	return _n
} // Name()

// Email returns the e-mail addrdd of the current git user, or the empty
// string if no e-mail address is configured.
func (u *user) Email() string {
	// is the git user e-mail defined in the environment?
	_e := os.Getenv("GIT_AUTHOR_EMAIL")
	if _e == "" {
		_e = os.Getenv("GIT_COMMITTER_EMAIL")
		if _e == "" {
			_e = u.email
			if _e == "" {
				_e = os.Getenv("EMAIL")
			}
		}
	}

	return _e
} // Email()

// String() returns a string representation of the git user's name and
// e-mail address, or the empty string if neither are defined.
func (u *user) String() string {
	_name := u.Name()
	_email := u.Email()

	if _name == "" {
		return _email
	} else if _email == "" {
		return _name
	} else {
		return fmt.Sprintf("%s <%s>", _name, _email)
	}
} // String()
