package main

import (
	"fmt"
	"strings"
)

func tag(long bool) string {
	_tag := fmt.Sprintf("%s v%s", exe(), VERSION)

	// do we need to display long version information?
	if long {
		_git := info()
		if _git != nil {
			// do we have a branch name?
			_branch, _ := _git.Branch()
			_strings := []string{}
			if _branch != "" {
				_strings = append(_strings, _branch)
			}

			// do we have a commit hash?
			_commit, _ := _git.Commit()
			if _commit != nil {
				if _commit.String() != "" {
					_strings = append(_strings, _commit.Prefix(BUILD))
				}
			}

			// combine the branch and commit
			_output := strings.Join(_strings, "/")

			// is this checkout modified?
			if _output != "" {
				_modified, _ := _git.Modified()
				if _modified {
					_output = _output + "+"
				}

				_tag = _tag + " " + _output
			}
		}
	}

	return _tag
} // version()
