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
			_strings := []string{_tag}
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

			// is this checkout modified?
			if len(_strings) > 1 {
				_modified, _ := _git.Modified()
				if _modified {
					_last := len(_strings) - 1
					_strings[_last] = _strings[_last] + "+"
				}
			}

			// extract the git version
			_version, _ := _git.Git()
			if _version != "" {
				_strings = append(_strings, "git", "v"+_version)
			}

			// augment the tag
			_tag = strings.Join(_strings, " ")
		}
	}

	return _tag
} // version()
