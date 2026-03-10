package v2

import "sync/atomic"

var (
	djinniTextIndex   int64
	djinniButtonIndex int64
)

func djinniTextLongOrShort() bool {
	return atomic.AddInt64(&djinniTextIndex, 1)&1 == 0
}

func djinniButtonSignUpOrJobs() bool {
	return atomic.AddInt64(&djinniButtonIndex, 1)&1 == 0
}
