package tests

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var host string

func TestSuite(t *testing.T) {
	host = os.Getenv("HOST")
	RegisterFailHandler(Fail)
	RunSpecs(t, fmt.Sprintf("%v Tests Suite", env))
}

type environment int

const (
	PROD environment = iota
	STAGING
	DEV
	LOCAL
)

func (e environment) String() string {
	return [...]string{"PROD", "STAGING", "DEV", "LOCAL"}[e]
}

// SkipIn takes 1 or more environment parameters if the current environment matches the supplied environments the test will not be executed.
// This method should be placed at the beginning of any test that needs to be excluded from running in specific environments
func SkipIn(environments ...environment){
	for _, e := range environments {
		if e == env {
			Skip(fmt.Sprintf("Skipped: %v, Env: %v",CurrentGinkgoTestDescription().TestText, e.String()))
		}
	}
}

// OnlyIn takes 1 or more environment parameters if the current environment does not any of the supplied environments the test will not be executed.
// This method should be placed at the beginning of any test that can only be run in specific environments
func OnlyIn(environments ...environment){
	for _, e := range environments {
		if e == env {
			return
		}
	}
	Skip(fmt.Sprintf("Skipped: %v, Env: %v",CurrentGinkgoTestDescription().TestText, env))
}


