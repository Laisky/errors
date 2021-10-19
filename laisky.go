package errors

import "sync/atomic"

var skipCallers int32 = 3

// SetSkipCallers set depth of skip callers
func SetSkipCallers(skip int) {
	atomic.StoreInt32(&skipCallers, int32(skip))
}

// GetSkipCallers get depth of skip callers
func GetSkipCallers() int {
	return int(atomic.LoadInt32(&skipCallers))
}
