package gitinfo

import (
	"runtime"
)

// Here returns the GitInfo instance referencing the caller's location. If the
// caller cannot be determined, Here returns the UnknownCallerError.
func Here() (GitInfo, error) {
	_, _file, _, _ok := runtime.Caller(1)
	if _ok {
		return NewWithPath(_file)
	} else {
		return nil, UnknownCallerError
	}
} // Here()
