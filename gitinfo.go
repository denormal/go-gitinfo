package gitinfo

import (
	"os"
	"strconv"
	"strings"

	"github.com/denormal/go-gitconfig"
	"github.com/denormal/go-gittools"
)

// the default git editor as detailed here:
//		  https://git-scm.com/book/en/v2/Customizing-Git-Git-Configuration
const _EDITOR = "vi"

// GitInfo represents basic information about a git working copy.
type GitInfo interface {
	// Branch returns the current branch name for the working copy. If the
	// GitInfo instance was initialised for a path not within a working copy,
	// Branch will return the empty string. An error is returned if there is
	// a problem determining the branch name.
	Branch() (string, error)

	// Commit returns the most recent Commit details for the working
	// copy. If the GitInfo instance was initialised for a path not within a
	// working copy, Commit will return nil. An error is returned if there is
	// a problem determining the commit details.
	Commit() (Commit, error)

	// Config returns the git configuration details for the working copy.
	// see https://github.com/denormal/go-gitconfig for more details.
	Config() gitconfig.GitConfig

	// Editor returns the git editor configured for working copy.
	Editor() string

	// Modified returns true if the working copy has been modified, either
	// through locally made changes, or untracked files. Modified returns an
	// error if a problem is encountered determining the modified state.
	Modified() (bool, error)

	// Path returns the absolute path used to initialised this GitInfo.
	Path() string

	// Root returns the root directory of the working copy. If the GitInfo
	// instance was initialised for a path not within a working copy, Root
	// returns the empty string.
	Root() string

	// User returns details of the git user for this working copy.
	User() User

	// Version returns the version string for the installed git executable,
	// or an error if this cannot be determined.
	Version() (string, error)

	// Map returns the git information as a map of strings.
	Map() map[string]string
}

type gitinfo struct {
	config gitconfig.GitConfig
}

// Config returns the git configuration details for the working copy.
// see https://github.com/denormal/go-gitconfig for more details.
func (g *gitinfo) Config() gitconfig.GitConfig { return g.config }

// Path returns the absolute path used to initialised this GitInfo.
func (g *gitinfo) Path() string { return g.config.Path() }

// Root returns the root directory of the working copy. If the GitInfo
// instance was initialised for a path not within a working copy, Root
// returns the empty string.
func (g *gitinfo) Root() string { return g.config.Root() }

// Commit returns the most recent Commit details for the working
// copy. If the GitInfo instance was initialised for a path not within a
// working copy, Commit will return nil. An error is returned if there is
// a problem determining the commit details.
func (g *gitinfo) Commit() (Commit, error) {
	// do we have a working copy root?
	_root := g.Root()
	if _root == "" {
		return nil, nil
	}

	// attempt to retrieve the current HEAD commit
	_bytes, _err := revparse(_root, "HEAD")
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

// Branch returns the current branch name for the working copy. If the GitInfo
// instance was initialised for a path not within a working copy, Branch will
// return the empty string. An error is returned if there is a problem
// determining the branch name.
func (g *gitinfo) Branch() (string, error) {
	// do we have a working copy root?
	_root := g.Root()
	if _root == "" {
		return "", nil
	}

	// attempt to retrieve the current HEAD commit
	_bytes, _err := revparse(_root, "--abbrev-ref", "HEAD")
	if _err != nil {
		if _err == MissingWorkingCopyError {
			return "", nil
		} else {
			return "", _err
		}
	}

	// extract the branch name
	return strings.TrimSpace(string(_bytes)), nil
} // Branch()

// Editor returns the git editor configured for working copy.
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

// Modified returns true if the working copy has been modified, either
// through locally made changes, or untracked files. Modified returns an
// error if a problem is encountered determining the modified state.
func (g *gitinfo) Modified() (bool, error) {
	// if we don't have a working copy root, then we can't determine
	// the modified status
	_root := g.Root()
	if _root == "" {
		return false, MissingWorkingCopyError
	}

	// attempt to determine the modified status
	_output, _err := gittools.RunInPath(_root, "status", "--porcelain")
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

// User returns details of the git user for this working copy.
func (g *gitinfo) User() User {
	return newUser(g.config)
} // User()

// Version returns the version string for the installed git executable,
// or an error if this cannot be determined.
func (g *gitinfo) Version() (string, error) {
	return gittools.Version()
} // Version()

// Map returns the git information as a map of strings.
func (g *gitinfo) Map() map[string]string {
	// extract the git information
	var (
		_branch, _   = g.Branch()
		_commit, _   = g.Commit()
		_modified, _ = g.Modified()
		_user        = g.User()
		_version, _  = g.Version()
	)

	// build the map
	_map := map[string]string{
		BRANCH:     _branch,
		EDITOR:     g.Editor(),
		PATH:       g.Path(),
		ROOT:       g.Root(),
		MODIFIED:   strconv.FormatBool(_modified),
		USER_NAME:  _user.Name(),
		USER_EMAIL: _user.Email(),
		VERSION:    _version,
	}

	// add the commit hash (if known)
	//		- ensure the COMMIT field is created, even if we don't have
	//		  a value
	//		- this ensures the map always contains all possible fields
	if _commit != nil {
		_map[COMMIT] = _commit.String()
	} else {
		_map[COMMIT] = ""
	}

	return _map
} // Map()

// ensure gitinfo supports the GitInfo interface
var _ GitInfo = &gitinfo{}
