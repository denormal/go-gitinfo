package gitinfo

import (
	"errors"

	"github.com/denormal/go-gittools"
)

var (
	MissingGitError         = gittools.MissingGitError
	MissingWorkingCopyError = gittools.MissingWorkingCopyError
	UnknownCallerError      = errors.New("unable to determine caller")
)
