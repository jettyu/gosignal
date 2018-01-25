package gosignal

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

var (
	// SIGUSR1 ...
	SIGUSR1 = syscall.Signal(10)
	// SIGUSR2 ...
	SIGUSR2 = syscall.Signal(12)
	// CURPID ...
	CURPID = 0
)

// SignalHandler ...
type SignalHandler struct {
	signalMap map[os.Signal]func()
	signals   chan os.Signal
}

// NewSignalHandler ...
func NewSignalHandler() SignalHandler {
	return SignalHandler{
		signalMap: make(map[os.Signal]func()),
		signals:   make(chan os.Signal),
	}
}

// Serve ...
func (p SignalHandler) Serve(pidfile ...string) {
	CURPID = os.Getpid()
	file := "pid"
	if len(pidfile) > 0 {
		file = pidfile[0]
	}
	ioutil.WriteFile(file, []byte(fmt.Sprint(CURPID)), os.ModePerm)
	for k := range p.signalMap {
		signal.Notify(p.signals, k)
	}
	for sig := range p.signals {
		callback, ok := p.signalMap[sig]
		if ok {
			callback()
		}
	}
}

// Signal ...
func (p SignalHandler) Signal(sig int, callback func()) {
	p.signalMap[syscall.Signal(sig)] = callback
}

var (
	signalHandler = NewSignalHandler()
)

// Signal ...
func Signal(sig int, callback func()) {
	signalHandler.Signal(sig, callback)
}

// Serve ...
func Serve(pidfile ...string) {
	signalHandler.Serve(pidfile...)
}
