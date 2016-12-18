package main

import (
	"fmt"
	"strings"

	"github.com/denormal/go-gitinfo"
)

func build(gi gitinfo.GitInfo, fields []string) (map[string]string, error) {
	_map := gi.Map()

	// should we only consider certain fields?
	var _rtn map[string]string
	if fields != nil {
		_rtn = make(map[string]string)
		for _, _field := range fields {
			_value, _ok := _map[_field]
			if _ok {
				_rtn[_field] = _value
			} else if strings.HasSuffix(_field, ".*") {
				_prefix := strings.TrimSuffix(_field, "*")
				_match := false
				for _f, _v := range _map {
					if strings.HasPrefix(_f, _prefix) {
						_rtn[_f] = _v
						_match = true
					}
				}
				if !_match {
					return nil, fmt.Errorf("no match for pattern %q", _field)
				}
			} else {
				return nil, fmt.Errorf("unknown field %q", _field)
			}
		}
	} else {
		_rtn = _map
	}

	// return the map
	return _rtn, nil
} // build()
