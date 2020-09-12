package tests

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

const (
	globalValue = "globalVal"
)
