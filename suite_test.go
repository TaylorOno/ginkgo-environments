package tests

import (
	"fmt"
	"github.com/TaylorOno/ginkgo-environments/config"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var host string

func TestSuite(t *testing.T) {
	host = os.Getenv("HOST")
	config.LoadKey()
	config.Encrypt()
	RegisterFailHandler(Fail)
	RunSpecs(t, fmt.Sprintf("%v Tests Suite", config.Env))
}