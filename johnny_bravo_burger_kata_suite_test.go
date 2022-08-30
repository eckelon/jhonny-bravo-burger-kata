package jhonny_bravo_burger_kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJhonnyBravoBurgerKata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JhonnyBravoBurgerKata Suite")
}
