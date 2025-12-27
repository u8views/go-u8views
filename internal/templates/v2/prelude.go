package v2

import "sync/atomic"

var djinniIndex int64

func djinniSignUpOrJobs() bool {
	return atomic.AddInt64(&djinniIndex, 1)&1 == 0
}
