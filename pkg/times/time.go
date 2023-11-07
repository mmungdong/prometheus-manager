package times

import (
	"time"
)

type UnixTime interface {
	Unix() int64
}

// TimeToUnix convert time.Time to unix timestamp
func TimeToUnix(t UnixTime) int64 {
	if t == nil {
		return 0
	}

	return t.Unix()
}

// UnixToTime convert unix timestamp to time.Time
func UnixToTime(unix int64) time.Time {
	return time.Unix(unix, 0)
}
