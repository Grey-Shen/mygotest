package exitutils

import (
	"fmt"
	"os"

	"github.com/onsi/gocleanup"
	"gitlab.qiniu.io/pingan/libserver/colorful"
)

func Success() {
	gocleanup.Exit(0)
}

func Failuref(format string, messages ...interface{}) {
	message := fmt.Sprintf(format, messages...)
	fmt.Fprintf(os.Stderr, colorful.StderrColor.Red(message))
	gocleanup.Exit(1)
}

func Failureln(message interface{}) {
	fmt.Fprintln(os.Stderr, colorful.StderrColor.Red(message))
	gocleanup.Exit(1)
}

func Failure() {
	gocleanup.Exit(1)
}
