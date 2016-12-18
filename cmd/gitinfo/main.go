package main

import (
	"flag"
	"os"
	"sort"
	"strings"

	"github.com/denormal/go-gitinfo"
)

// define the command version
const (
	VERSION = "0.01"
	BUILD   = 10
)

type options struct {
	env     *bool   // environment only: editor,user.*,path,root,version
	fields  *string // explicit list of fields
	h       *bool   // short help
	help    *bool   // full help
	output  *string // output to this file
	r       *bool   // runtime update of the package symbol
	runtime *bool   //		- as with 'r'
	s       *bool   // short output without field names
	short   *bool   //		- as with 's'
	src     *bool   // source information only: commit,branch,modified
	symbol  *string // the package symbol
	v       *bool   // output short version information
	version *bool   // output detailed version information
}

var opt *options

func main() {
	// handle the command line arguments
	flag.Parse()
	if *opt.help {
		usage(true)
	} else if *opt.h {
		usage(false)
	} else if *opt.version {
		version(true)
	} else if *opt.v {
		version(false)
	}

	// are we outputting to a file or stdout?
	var (
		_out = os.Stdout
		_err error
	)
	if *opt.output != "" {
		_out, _err = os.Create(*opt.output)
		if _err != nil {
			fail(1,
				"%s: error: %s: %s\n",
				exe(), *opt.output, _err.Error(),
			)
		}
		defer _out.Close()
	}

	// should we only output certain fields?
	var _f []string
	if *opt.fields != "" {
		_f = strings.Split(*opt.fields, ",")
	} else if *opt.env {
		_f = []string{
			gitinfo.EDITOR,
			gitinfo.PATH,
			gitinfo.ROOT,
			gitinfo.USER_EMAIL,
			gitinfo.USER_NAME,
			gitinfo.VERSION,
		}
	} else if *opt.src {
		_f = []string{
			gitinfo.BRANCH,
			gitinfo.COMMIT,
			gitinfo.MODIFIED,
		}
	}

	// have we been given a symbol?
	//		- is it well-formed?
	var (
		_pkg string
		_var string
	)
	if *opt.symbol != "" {
		_parts := strings.Split(*opt.symbol, ".")
		if len(_parts) != 2 ||
			_parts[0] == "" ||
			_parts[1] == "" {
			fail(1,
				"%s: invalid symbol %q; expected %q\n",
				exe(), *opt.symbol, "pkg.var",
			)
		}
		_pkg, _var = _parts[0], _parts[1]
	}

	// have we been given a path?
	//		- attempt to load the gitinfo for this path or the current path
	var (
		_info gitinfo.GitInfo
	)
	if len(flag.Args()) == 0 {
		_info, _err = gitinfo.New()
	} else {
		_info, _err = gitinfo.NewWithPath(flag.Arg(0))
	}

	// did we encounter an error?
	if _err != nil {
		fail(2, "%s: error: %s\n", exe(), _err.Error())
	} else if _info != nil {
		_map, _err := build(_info, _f)
		if _err != nil {
			fail(3, "%s: error: %s\n", exe(), _err.Error())
		} else if _pkg == "" {
			// do we want a short or long display?
			//		- i.e. just the values, or key = value?
			display(_out, _map, *opt.s || *opt.short, _f)
		} else {
			generate(_out, _map, _pkg, _var, *opt.r || *opt.runtime)
		}
	}

	// everything is OK
	ok()
} // main()

func init() {
	_b := func(f, msg string) *bool { return flag.Bool(f, false, msg) }
	_s := func(f, msg string) *string { return flag.String(f, "", msg) }

	// extract the list of supported field names
	_fields := []string{}
	_seen := make(map[string]bool)
	for _f, _ := range gitinfo.Build(nil).Map() {
		_fields = append(_fields, _f)
		if strings.Contains(_f, ".") {
			_parts := strings.SplitN(_f, ".", 2)
			if !_seen[_parts[0]] {
				_fields = append(_fields, _parts[0]+".*")
				_seen[_parts[0]] = true
			}
		}
	}
	sort.Strings(_fields)

	// define the fields help
	_text := []string{}
	for _, _field := range _fields {
		_text = append(_text, "\t\t"+_field)
	}

	// define the command flags
	opt = &options{
		h:       _b("h", "Display short help."),
		help:    _b("help", "Display detailed command help."),
		v:       _b("v", "Display the command version."),
		version: _b("version", "Display the detailed command version."),
		r:       _b("r", ""),
		runtime: _b("runtime",
			"When used with -X, allow runtime checking of the git "+
				"information,\n"+
				"\tfalling back to the compiled values set by -X if the git "+
				"information\n"+
				"\tcannot be determined.",
		),
		s:     _b("s", ""),
		short: _b("short", "Short display; only output field values."),

		env: _b("env",
			"Environment information only; equivalent to\n"+
				"\t-f editor,path,root,user.*,version.",
		),
		src: _b("src",
			"Source information only; equivalent to "+
				"-f branch,commit,modified.",
		),

		fields: _s("f",
			"Output just the given `fields` (comma-separated); choose from:\n"+
				strings.Join(_text, "\n"),
		),
		output: _s("o", "Output to `path` instead of STDOUT."),
		symbol: _s("X",
			"Output the git information to the package variable `pkg.var` "+
				"of type \n"+
				"\tgitinfo.GitInfo. This is intended to be used with "+
				"go:generate.",
		),
	}
} // init()
