// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"strings"
)

// Parse returns an argument map from a parameter slice and argument slice.
func Parse(paras, elems []string) (map[string]string, error) {
	amap := make(map[string]string)

	for i, para := range paras {
		name, dflt, ok := strings.Cut(para, ":")
		switch {
		case i >= len(elems) && ok:
			amap[name] = dflt
		case i >= len(elems) && !ok:
			return nil, fmt.Errorf("cannot parse argument %q - not provided", name)
		default:
			amap[name] = elems[i]
		}
	}

	return amap, nil
}

// Split returns a Command name and argument slice from an argument slice.
func Split(elems []string) (string, []string) {
	switch len(elems) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(elems[0]), nil
	default:
		return strings.ToLower(elems[0]), elems[1:]
	}
}
