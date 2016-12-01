package gitinfo

import (
	"github.com/denormal/go-gittools"
)

func revparse(path string, args ...string) ([]byte, error) {
	// ensure we are in a git working copy
	_is, _err := gittools.IsWorkingCopy(path)
	if _err != nil {
		return nil, _err
	} else if !_is {
		return nil, gittools.MissingWorkingCopyError
	}

	// append the given arguments to the "rev-parse" command
	_args := append([]string{"rev-parse"}, args...)

	return gittools.RunInPath(path, _args...)
} // revparse()
