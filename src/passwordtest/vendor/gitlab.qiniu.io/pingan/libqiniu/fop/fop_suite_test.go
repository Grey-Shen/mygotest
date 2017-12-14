package fop_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFop(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fop Suite")
}
