package pack

import (
	"fmt"
	"os"
)

var (
	Version     = "v0.0.0"
	Build_at    = "2016-11-18 16:40:00"
	Commit_sha1 = ""
	Go_version  = ""
)

func ShowVersion() {
	fmt.Printf("version: %s\nbuild: %s\ncommit: %s\ngo: %s\n", Version, Build_at, Commit_sha1, Go_version)
	os.Exit(0)
}
