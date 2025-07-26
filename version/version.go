package version

import "fmt"

var gVersion = "1.0.0"
var gBuildTime = ""
var gCommitID = ""

func VersionInfo() {
	fmt.Printf("Version:%s\n", gVersion)

	if gBuildTime != "" {
		fmt.Printf("BuildTime:%s\n", gBuildTime)
	}

	if gCommitID != "" {
		fmt.Printf("CommitID:%s\n", gCommitID)
	}
}

func Version() string {
	return gVersion
}

func BuildTime() string {
	return gBuildTime
}

func CommitID() string {
	return gCommitID
}
