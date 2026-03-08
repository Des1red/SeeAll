package devmode

import "sync/atomic"

var isPrewarming atomic.Bool

func SetPrewarming(v bool) {
	isPrewarming.Store(v)
}

func IsPrewarming() bool {
	return isPrewarming.Load()
}
