package env

import (
	"fmt"
	"strings"
)

const (
	Release = "release"
	Debug   = "debug"
	Test    = "test"
)

var env string

func Set(e string) error {
	if err := checkEnv(e); err != nil {
		return err
	}
	env = e
	return nil
}

func Get() string {
	return env
}

func checkEnv(e string) error {
	e = strings.ToLower(e)
	switch e {
	case Release, Test, Debug:
		return nil
	default:
		return fmt.Errorf("Invalid Environment: %s, only `debug`, `release` and `test` are supported.", e)
	}
}
