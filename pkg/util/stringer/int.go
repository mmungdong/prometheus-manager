package stringer

import (
	"strconv"
)

// ParseInt32 parse int32
func ParseInt32(s string) int32 {
	return int32(ParseInt64(s))
}

// ParseInt64 parse int64
func ParseInt64(s string) int64 {
	v, _ := ParseInt(s)
	return v
}

// ParseInt parse int
func ParseInt(s string) (int64, error) {
	return ParseIntWithDefault(s, 0)
}

// ParseIntWithDefault parse int with default value
func ParseIntWithDefault(s string, def int64) (int64, error) {
	if s == "" {
		return def, nil
	}
	return ParseIntWithError(s)
}

// ParseIntWithError parse int with error
func ParseIntWithError(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
