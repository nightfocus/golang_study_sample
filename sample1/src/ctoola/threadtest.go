// +build linux

package ctoola

import (
	"golang.org/x/sys/unix"
)

func GetThreadID() uint64 {
	tid := unix.Gettid()

	return (uint64)(tid)
}
