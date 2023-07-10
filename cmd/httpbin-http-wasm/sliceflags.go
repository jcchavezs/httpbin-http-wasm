package main

import (
	"flag"
	"strings"
)

type sliceFlags []string

var _ flag.Value = (*sliceFlags)(nil)

func (i *sliceFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *sliceFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
