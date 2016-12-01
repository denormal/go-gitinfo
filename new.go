package gitinfo

import (
	"os"

	"github.com/denormal/go-gitconfig"
	"github.com/denormal/go-gittools"
)

func New() (GitInfo, error) {
	return NewWithPath("")
} // New()

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

	// attempt to extract the working copy root
	_root, _err := gittools.WorkingCopy(path)
	if _err != nil {
		if _err != gittools.MissingWorkingCopyError {
			return nil, _err
		}
	}

	// create the GitInfo instance
	_info := &gitinfo{
		config: _config,
		root:   _root,
	}

	return _info, nil
} // NewWithPath()
