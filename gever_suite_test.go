package gever_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGever(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gever Suite")
}
