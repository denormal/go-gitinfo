# go-gitinfo

Package `gitinfo` provides an interface to high-level git information for a
git working copy. `gitinfo` attempts to use the locally installed
`git` executable using [go-gittools](https://github.com/denormal/go-gittools).

```go
import "github.com/denormal/go-gitinfo"

// load git information for a particular working copy
info, err := gitinfo.NewWithPath("/my/git/working/copy")
if err != nil {
    panic(err)
}

// is this working copy modified?
modified, err := info.Modified()
if err != nil {
    panic(err)
} else if modified {
    fmt.Printf("%s is modified\n", info.Root())
}

// what is the head commit hash?
commit, err := info.Commit()
if err != nil {
    panic(err)
} else if commit != nil {
    fmt.Printf("%s HEAD is at commit %s\n", info.Root(), commit.Prefix(8))
}
```

For more information see `godoc github.com/denormal/go-gitinfo`.

## Installation

`go-gitinfo` can be installed using the standard Go approach:

```sh
go install github.com/denormal/go-gitinfo
```

`go-gitinfo` also provides a command-line tool for extracting git information:

```sh
go install github.com/denormal/go-gitinfo/cmd/gitinfo
```

Command help may be accessed using
```sh
% gitinfo --help
```

## License

Copyright (c) 2016 Denormal Limited

[MIT License](LICENSE)
