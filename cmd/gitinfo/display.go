package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func display(out io.Writer, m map[string]string, short bool, fields []string) {
	// have we been given a list of fields to display?
	//		- if so, we use the given order
	if fields == nil {
		// sort the map keys
		//		- determine the maximum key length
		fields = make([]string, 0, len(m))
		for _k, _ := range m {
			fields = append(fields, _k)
		}
		sort.Strings(fields)
	}

	// do we have any fields that are wildcard patterns (e.g. user.*)
	//		- expand them to all matched fields
	_fields := make([]string, 0)
	for _, _f := range fields {
		if strings.HasSuffix(_f, ".*") {
			_prefix := strings.TrimSuffix(_f, "*")
			_match := make([]string, 0)
			for _k, _ := range m {
				if strings.HasPrefix(_k, _prefix) {
					_match = append(_match, _k)
				}
			}
			sort.Strings(_match)
			_fields = append(_fields, _match...)
		} else {
			_fields = append(_fields, _f)
		}
	}

	// if we are displaying the field name in the output, determine the
	// maximum length of the name to support left-justified output
	_len := 0
	if !short {
		for _, _f := range _fields {
			_l := len(_f)
			if _l > _len {
				_len = _l
			}
		}
	}

	// output the map entries in field order
	for _, _f := range _fields {
		if short {
			fmt.Fprintf(out, "%s\n", m[_f])
		} else {
			fmt.Fprintf(out, "%-*s = %s\n", _len, _f, m[_f])
		}
	}
} // display()
