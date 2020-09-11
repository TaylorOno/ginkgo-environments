package tests

import (
	"fmt"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Functional", func() {
	Context("Tests", func(){
		It("Prints system variable", func(){
			fmt.Println("this is a system variable: " + host)
		})

		It("Prints global variable", func(){
			fmt.Println("this is a global variable: " + globalValue)
		})

		It("Prints environment specific variable", func(){
			fmt.Println("this is a env variable: " + user)
		})

		It("Skips Destructive Tests in PROD", func(){
			SkipIn(PROD)
			if env == PROD {
				Fail(fmt.Sprintf("This will fail in prod"))
			}
			fmt.Println("this does not run in prod")
		})

		It("Runs Destructive Tests only in Local", func(){
			OnlyIn(LOCAL)
			if env != LOCAL {
				Fail(fmt.Sprintf("This will fail if not local"))
			}
		})
	})
})
