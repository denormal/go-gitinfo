package main

import (
	"github.com/denormal/go-gitinfo"
)

var git gitinfo.GitInfo

func info() gitinfo.GitInfo {
	// if we don't have the git information already, attempt to load
	// it for the current location
	//		- we may use go:generate to compile this tool with a fixed
	//		  git information if we need to distribute the gitinfo tool
	//		  without source code
	if git == nil {
		git, _ = gitinfo.Here()
	}

	return git
} // info()

// run go:generate if gitinfo needs to be shipped without source code
//		- this will capture its git information as a compile-time structure
//
//go:generate gitinfo -o git.go -runtime -src -X main.git
