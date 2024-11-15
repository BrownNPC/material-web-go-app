package md

import (
	_ "embed"
)

//go:embed init.html
var headers string

func RawHeaders() []string {
	return []string{headers}
}
