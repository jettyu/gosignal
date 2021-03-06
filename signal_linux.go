//+build linux

package gosignal

import (
	"syscall"
)

func init() {
	SIGUSR1 = syscall.SIGUSR1
	SIGUSR2 = syscall.SIGUSR2
}

// Kill ...
func Kill(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}
