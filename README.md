# Ginkgo Multi Environment Test Setup
This is a sample solution that will allow ginkgo tests to be easily run against different environments

## Getting Started
Global configs should be placed in the `config.go` test file.  These are variables that can be shared across all environments.
Environment specific variables can should be placed in the appropriate `config_{env).go` test file.

## Helper Methods
There are two helper methods designed to help you conditionally skip test `OnlyIn()` and `SkipIn()` these are variadic functions that take zero or more environments as argument.
Placing these one of these methods at the beginning of your test will help you prevent potentially destructive tests running in the wrong environment.
For Exampled `OnlyIn(LOCAL,DEV)` will cause a test to be skipped in PROD and STAGING.  Conversely `SkipIn(PROD, STAGING)` will have the same effect.

## Example execution
### Run Test Local
This can be accomplished using the default ginkgo command `ginkgo` as Local is the default profile
```
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
this can be accomplished passing `dev` as a tag for example `ginkgo --tags dev`
```
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
  
3) Updates the environment enum in config.go
    - add your {env} to the var block
    - be careful to also create a simple name for the String Method.