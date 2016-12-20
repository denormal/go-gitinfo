package gitinfo

import (
	"strconv"

	"github.com/denormal/go-gitconfig"
)

const (
	BRANCH     = "branch"
	COMMIT     = "commit"
	EDITOR     = "editor"
	GIT        = "git"
	MODIFIED   = "modified"
	PATH       = "path"
	ROOT       = "root"
	USER_NAME  = "user.name"
	USER_EMAIL = "user.email"
)

func Build(kv map[string]string) GitInfo {
	// is the modified flag set?
	_modified := false
	if kv[MODIFIED] == "true" {
		_modified = true
	}

	// return the GitInfo structure
	return &build{
		gitinfo:  gitinfo{},
		branch:   kv[BRANCH],
		commit:   &commit{kv[COMMIT]},
		editor:   kv[EDITOR],
		git:      kv[GIT],
		modified: _modified,
		path:     kv[PATH],
		root:     kv[ROOT],
		user:     &user{kv[USER_NAME], kv[USER_EMAIL]},
	}
} // Build()

type build struct {
	gitinfo

	branch   string
	commit   Commit
	editor   string
	git      string
	modified bool
	path     string
	root     string
	user     User
}

func (b build) Branch() (string, error)     { return b.branch, nil }
func (b build) Commit() (Commit, error)     { return b.commit, nil }
func (b build) Config() gitconfig.GitConfig { return nil }
func (b build) Path() string                { return b.path }
func (b build) Root() string                { return b.root }
func (b build) Editor() string              { return b.editor }
func (b build) Modified() (bool, error)     { return b.modified, nil }
func (b build) User() User                  { return b.user }
func (b build) Git() (string, error)        { return b.git, nil }

func (b build) Map() map[string]string {
	return map[string]string{
		BRANCH:     b.branch,
		COMMIT:     b.commit.String(),
		EDITOR:     b.editor,
		GIT:        b.git,
		MODIFIED:   strconv.FormatBool(b.modified),
		PATH:       b.path,
		ROOT:       b.root,
		USER_EMAIL: b.user.Email(),
		USER_NAME:  b.user.Name(),
	}
} // Map()

// ensure the static type implements the GitInfo interface
var _ GitInfo = &build{}
