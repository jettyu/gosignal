package gosignal

import (
	"syscall"
)

//+build drwin

// Kill ...
func Kill(pid int, sid syscall.Signal) error {
	return nil
}
