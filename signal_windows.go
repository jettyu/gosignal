//+build windows

package gosignal

import (
	"syscall"
)

// Kill ...
func Kill(pid int, sid syscall.Signal) error {
	return nil
}
