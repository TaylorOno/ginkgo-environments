package tests

import (
	"fmt"
	"github.com/TaylorOno/ginkgo-environments/config"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Functional", func() {
	Context("Tests", func(){
		It("Prints system variable", func(){
			fmt.Println("this is a system variable: " + host)
		})

		It("Prints global variable", func(){
			fmt.Println("this is a global variable: " + config.GlobalValue)
		})

		It("Prints environment specific variable", func(){
			fmt.Println("this is a env variable: " + config.User)
		})

		It("Prints secret specific variable", func(){
			fmt.Println("this is a password: " + config.Password.Decrypt())
		})

		It("Skips Destructive Tests in PROD", func(){
			config.SkipIn(config.PROD)
			if config.Env == config.PROD {
				Fail(fmt.Sprintf("This will fail in prod"))
			}
			fmt.Println("this does not run in prod")
		})

		It("Runs Destructive Tests only in Local", func(){
			config.OnlyIn(config.LOCAL)
			if config.Env != config.LOCAL {
				Fail(fmt.Sprintf("This will fail if not local"))
			}
		})
	})
})
