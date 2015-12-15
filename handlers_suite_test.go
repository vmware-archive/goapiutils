package goapiutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApiUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Utils Suite")
}
