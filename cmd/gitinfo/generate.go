package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"time"

	"go/format"

	"github.com/denormal/go-gitinfo"
)

func generate(out io.Writer, m map[string]string, pkg, v string, runtime bool) {
	// extract the command line argument used in this invocation
	_cmd := strings.Join(os.Args, " ")

	// generate a string representation of the map
	//		- we use ordered strings to ensure repeatability
	_keys := make([]string, 0, len(m))
	for _k, _ := range m {
		_keys = append(_keys, _k)
	}
	sort.Strings(_keys)

	_map := ""
	for _, _k := range _keys {
		_map = _map + fmt.Sprintf("%q : %q,\n", _k, m[_k])
	}

	// determine the import path of go-gitinfo
	//		- we could hard-code this, but instead we use reflection
	_ref := reflect.TypeOf((*gitinfo.GitInfo)(nil)).Elem()
	_import := _ref.PkgPath()

	// should we include runtime checking of the git information?
	_runtime := ""
	if runtime {
		_runtime = fmt.Sprintf("%s, _ = gitinfo.Here()", v)
	}

	// generate the package definition
	_src := fmt.Sprintf(
		`
// generated by %s
//           on %s
//   DO NOT EDIT; local changes will be overridden
//
//   %% %s
package %s

import %q

func init() { %s
    if %s == nil {
        %s = gitinfo.Build( map[string]string{
                %s
             } )
    }
}`,
		tag(true), time.Now().UTC().String(), _cmd,
		pkg, _import,
		_runtime, v, v, _map,
	)

	// apply standard formatting
	_bytes, _err := format.Source([]byte(_src))
	if _err != nil {
		panic(_err)
	}

	fmt.Fprint(out, string(_bytes))
} // generate()
