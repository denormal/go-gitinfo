package gitinfo

import (
	"os"

	"github.com/denormal/go-gitconfig"
)

// New returns the GitInfo instance for the current process working directory,
// or an error if the directory cannot be resolved, or the git executable
// cannot be found.
func New() (GitInfo, error) {
	return NewWithPath("")
} // New()

// NewWithPath returns the GitInfo instance for the given path, or an error
// if the path cannot be resolved or the git executable cannot be found. If
// path is "", NewWithPath examines the current process working directory.
func NewWithPath(path string) (GitInfo, error) {
	var _err error

	// if we have an empty path, then choose the current working directory
	if path == "" {
		path, _err = os.Getwd()
		if _err != nil {
			return nil, _err
		}
	}

	// attempt to load the git configuration
	_config, _err := gitconfig.NewWithPath(path)
	if _err != nil {
		return nil, _err
	}

	// create the GitInfo instance
	_info := &gitinfo{
		config: _config,
	}

	return _info, nil
} // NewWithPath()
