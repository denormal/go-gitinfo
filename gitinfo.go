package gitinfo

import (
	"os"
	"regexp"
	"strings"

	"github.com/denormal/go-gitconfig"
	"github.com/denormal/go-gittools"
)

// the default git editor as detailed here:
//		  https://git-scm.com/book/en/v2/Customizing-Git-Git-Configuration
const _EDITOR = "vi"

// declare the regular expression for matching version strings
var _VERSION *regexp.Regexp

type GitInfo interface {
	Commit() (Commit, error)
	Config() gitconfig.Config
	Editor() string
	Modified() (bool, error)
	Root() string
	User() User
	Version() (string, error)
}

type gitinfo struct {
	config gitconfig.Config
	root   string
}

func (g *gitinfo) Config() gitconfig.Config { return g.config }
func (g *gitinfo) Root() string             { return g.root }

func (g *gitinfo) Commit() (Commit, error) {
	// do we have a working copy root?
	if g.root == "" {
		return nil, nil
	}

	// attempt to retrieve the current HEAD commit
	_bytes, _err := revparse(g.root, "HEAD")
	if _err != nil {
		if _err == MissingWorkingCopyError {
			return nil, nil
		} else {
			return nil, _err
		}
	}

	// extract the commit string
	_commit := strings.TrimSpace(string(_bytes))
	if _commit == "" {
		return nil, nil
	}

	// return the commit instance
	return newCommit(_commit), nil
} // Commit()

func (g *gitinfo) Editor() string {
	// examine the environment for the editor
	for _, _env := range []string{"GIT_EDITOR", "EDITOR", "VISUAL"} {
		_editor := os.Getenv(_env)
		if _editor != "" {
			return _editor
		}
	}

	// attempt to extract the editor form the git configuration
	//		- if there's no configuration set, default to "vi"
	//		- this is consistent with
	//		  https://git-scm.com/book/en/v2/Customizing-Git-Git-Configuration
	_editor := g.config.Get("core.editor")
	if _editor == nil {
		return _EDITOR
	} else {
		return _editor.String()
	}
} // Editor()

func (g *gitinfo) Modified() (bool, error) {
	// if we don't have a working copy root, then we can't determine
	// the modified status
	if g.root == "" {
		return false, MissingWorkingCopyError
	}

	// attempt to determine the modified status
	_output, _err := gittools.RunInPath(g.root, "status", "--porcelain")
	if _err != nil {
		return false, _err
	}

	// do we have any non-empty lines?
	for _, _line := range strings.Split(string(_output), "\n") {
		_line = strings.TrimSpace(_line)
		if _line != "" {
			return true, nil
		}
	}

	return false, nil
} // Modified()

func (g *gitinfo) User() User {
	return newUser(g.config)
} // User()

func (g *gitinfo) Version() (string, error) {
	// attempt to extract the version of the current git executable
	//		- we don't need to be in a particular directory for this
	//		  so default to the current directory
	_bytes, _err := gittools.Run("--version")
	if _err != nil {
		return "", _err
	}

	// attempt to parse the version string
	//		- we're looking for a dotted sequence of numbers
	return _VERSION.FindString(string(_bytes)), nil
} // Version()

// ensure gitinfo supports the GitInfo interface
var _ GitInfo = &gitinfo{}

func init() {
	// compile the regular expression pattern
	//		- we're looking for a.c.b numbers
	//		- this may need to be expanded to include modifiers (e.g. "-dev")
	_VERSION = regexp.MustCompile("\\d+(\\.\\d+)*")
} // init()
