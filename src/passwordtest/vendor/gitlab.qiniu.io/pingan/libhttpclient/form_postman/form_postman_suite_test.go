package form_postman_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFormPostman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FormPostman Suite")
}
