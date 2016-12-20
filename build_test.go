package gitinfo_test

import (
	"testing"

	"github.com/denormal/go-gitinfo"
)

const _NONSENSE = "my.nonsense.field"

func TestBuild(t *testing.T) {
	// create a GitInfo instance
	_map := map[string]string{
		gitinfo.BRANCH:     "branch",
		gitinfo.COMMIT:     "commit",
		gitinfo.EDITOR:     "editor",
		gitinfo.GIT:        "git",
		gitinfo.MODIFIED:   "true",
		gitinfo.PATH:       "path",
		gitinfo.ROOT:       "root",
		gitinfo.USER_NAME:  "user.name",
		gitinfo.USER_EMAIL: "user.email",
		_NONSENSE:          "nonsense",
	}

	// ensure Build creates the requisite model
	_git := gitinfo.Build(_map)
	//		- branch
	_branch, _err := _git.Branch()
	if _err != nil {
		t.Fatalf("unexpected error in Branch(): %s", _err.Error())
	} else if _branch != _map[gitinfo.BRANCH] {
		t.Fatalf("unexpected Branch(); expected %q, got %q",
			_map[gitinfo.BRANCH], _branch,
		)
	}
	//		- config
	_config := _git.Config()
	if _config != nil {
		t.Fatalf(
			"unexpected Config(); expected %v, got %v", nil, _config,
		)
	}
	//		- commit
	_commit, _err := _git.Commit()
	if _err != nil {
		t.Fatalf("unexpected error in Commit(): %s", _err.Error())
	} else if _commit == nil {
		t.Fatalf("unexpected nil from Commit()")
	} else if _commit.String() != _map[gitinfo.COMMIT] {
		t.Fatalf(
			"unexpected Commit(); expected %q, got %q",
			_map[gitinfo.COMMIT], _commit.String(),
		)
	}
	//		- editor
	_editor := _git.Editor()
	if _editor != _map[gitinfo.EDITOR] {
		t.Fatalf(
			"unexpected Editor(); expected %q, got %q",
			_map[gitinfo.EDITOR], _editor,
		)
	}
	//		- modified
	_modified, _err := _git.Modified()
	if _err != nil {
		t.Fatalf("unexpected error in Modified(): %s", _err.Error())
	} else if !_modified { // the test set above uses "true"
		t.Fatalf(
			"unexpected Modified(); expected %s, got %v",
			_map[gitinfo.MODIFIED], _modified,
		)
	}
	//		- path
	_path := _git.Path()
	if _path != _map[gitinfo.PATH] {
		t.Fatalf(
			"unexpected Path(); expected %q, got %q",
			_map[gitinfo.PATH], _path,
		)
	}
	//		- root
	_root := _git.Root()
	if _root != _map[gitinfo.ROOT] {
		t.Fatalf(
			"unexpected Root(); expected %q, got %q",
			_map[gitinfo.ROOT], _root,
		)
	}
	//		- user
	_user := _git.User()
	if _user == nil {
		t.Fatalf("unexpected nil from User()")
	} else if _user.Name() != _map[gitinfo.USER_NAME] {
		t.Fatalf(
			"unexpected User() name; expected %q, got %q",
			_map[gitinfo.USER_NAME], _user.Name(),
		)
	} else if _user.Email() != _map[gitinfo.USER_EMAIL] {
		t.Fatalf(
			"unexpected User() name; expected %q, got %q",
			_map[gitinfo.USER_EMAIL], _user.Email(),
		)
	}
	//		- git version
	_gitversion, _err := _git.Git()
	if _err != nil {
		t.Fatalf("unexpected error in Git(): %s", _err.Error())
	} else if _gitversion != _map[gitinfo.GIT] {
		t.Fatalf(
			"unexpected Git(); expected %q, got %q",
			_map[gitinfo.GIT], _gitversion,
		)
	}

	// Map()
	_got := _git.Map()
	for _k, _v := range _got {
		_value, _ok := _map[_k]
		if !_ok {
			t.Fatalf("expected map value for %q, none found", _k)
		} else if _value != _v {
			t.Fatalf(
				"unexpected value for %q; expected %q, got %q",
				_k, _v, _value,
			)
		}
	}
	//		- ensure the "nonsense" key is not returned
	_nonsense, _ok := _got[_NONSENSE]
	if _ok {
		t.Fatal(
			"unexpected result from Map(); expected no entry for %q, got %q",
			_NONSENSE, _nonsense,
		)
	}
} // TestBuild()
