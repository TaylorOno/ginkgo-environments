# Ginkgo Multi Environment Test Setup
This is a sample solution that will allow ginkgo tests to be easily run against different environments

## Getting Started
Global configs should be placed in the `config-gobal.go` test file.  These are variables that can be shared across all environments.
Environment specific variables can should be placed in the appropriate `config-{env).go` test file.

## Helper Methods
There are two helper methods designed to help you conditionally skip test `OnlyIn()` and `SkipIn()` these are Variadic that take 0 or more environments as argument. 
Placing these in a test spec will cause a test to be only run in or be skipped in the specified environments.

## Example execution
### Run Test Local
this can be accomplished using the default ginkgo command as Local is the default profile
```bash
HOST=localhost:8443 ginkgo
Running Suite: LOCAL Tests Suite
================================
Random Seed: 1599804804
Will run 5 of 5 specs

this is a system variable: localhost:8443
•this is a global variable: globalVal
•this is a env variable: localUser
•this does not run in prod
••
Ran 5 of 5 Specs in 0.001 seconds
SUCCESS! -- 5 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

```

### Run Test Dev
this can be accomplished using the default ginkgo command as Local is the default profile
```bash
HOST=localhost:8443 ginkgo --tags dev
Running Suite: DEV Tests Suite
==============================
Random Seed: 1599804928
Will run 5 of 5 specs

this is a system variable: localhost:8443
•this is a global variable: globalVal
•this is a env variable: devUser
•this does not run in prod
•
------------------------------
S [SKIPPING] [0.000 seconds]
Functional
/home/taylor/workspace/http-cacheable-demo/tests/functional_test.go:8
  Tests
  /home/taylor/workspace/http-cacheable-demo/tests/functional_test.go:9
    Runs Destructive Tests only in Local [It]
    /home/taylor/workspace/http-cacheable-demo/tests/functional_test.go:30

    Skipped: Runs Destructive Tests only in Local, Env: DEV

    /home/taylor/workspace/http-cacheable-demo/tests/suite_test.go:47
------------------------------

Ran 4 of 5 Specs in 0.001 seconds
SUCCESS! -- 4 Passed | 0 Failed | 0 Pending | 1 Skipped
PASS
```

## Adding New Environments
assuming env is the desired name of the new environment

1) Create a new config_{env}.go file 
  - update the build tag to match the file env
  - updated any environment specific variables at the very least env={env}
  
2) Update the config_local.go file
  - add !{env} to the list of excluded build tags
  
3) Updates the environment enum in suite_test.go
  - add your {env} to the var block
  - be careful to also create a simple name for the String Method.