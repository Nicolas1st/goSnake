package routines

import (
	"os"
	"os/signal"
)

// CaptureSystemTerminationCall - relays the fact
// of system call occurence to another routine
func CaptureSystemTerminationCall(systemCallChannel chan bool) {
	// channel using which the os will send signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// the routine block until it receives a signal
	<-signalChannel

	// as soon as the signal above is received, this executes
	systemCallChannel <- true
}
