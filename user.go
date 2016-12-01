package gitinfo

import (
	"fmt"
	"os"

	"github.com/denormal/go-gitconfig"
)

type User interface {
	Name() string
	Email() string

	String() string
}

type user struct {
	name  string
	email string
}

func newUser(config gitconfig.Config) User {
	// extract the name & email from the git configuration
	_name := ""
	_config := config.Get("user.name")
	if _config != nil {
		_name = _config.String()
	}
	_email := ""
	_config = config.Get("user.email")
	if _config != nil {
		_email = _config.String()
	}

	return &user{name: _name, email: _email}
} // newUser()

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
