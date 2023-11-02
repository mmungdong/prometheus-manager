package model

import (
	"fmt"
)

type (
	Category int
	Status   int
)

const unknown = "未知"

func (s Status) String() string {
	switch s {
	case Enable:
		return "启用"
	default:
		return unknown
	}
}

func (c Category) String() string {
	switch c {
	default:
		return unknown
	}
}

var _ fmt.Stringer = (*Category)(nil)
var _ fmt.Stringer = (*Status)(nil)

const (
	Enable Status = iota + 1
)
