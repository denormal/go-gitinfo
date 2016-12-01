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
```

For more information see `godoc github.com/denormal/go-gitinfo`.

## Installation

`go-gitinfo` can be installed using the standard Go approach:

```go
go get github.com/denormal/go-gitinfo
```

## License

Copyright (c) 2016 Denormal Limited

[MIT License](LICENSE)
