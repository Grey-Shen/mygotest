package pidfile

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/onsi/gocleanup"
)

func Create(pidFilePath string) error {
	if pidFilePath == "" {
		return nil
	}
	pid := os.Getpid()
	if err := ioutil.WriteFile(pidFilePath, []byte(strconv.Itoa(pid)), 0600); err != nil {
		return err
	}
	gocleanup.Register(func() {
		os.Remove(pidFilePath)
	})
	return nil
}
