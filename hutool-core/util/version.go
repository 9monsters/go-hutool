package util

import "fmt"

var (
	GitCommit  = ""
	BuildTime  = ""
	GoVersion  = ""
	AppVersion = "1.0.0"
)

// Version returns the full version information for the application.
func Version() string {
	return fmt.Sprintf("version: %s\n", AppVersion) +
		fmt.Sprintf("build  : %s\n", BuildTime) +
		fmt.Sprintf("git    : %s\n", GitCommit) +
		fmt.Sprintf("go     : %s", GoVersion)
}
