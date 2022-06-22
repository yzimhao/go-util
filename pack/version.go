package pack

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func NewTag(version string) {
	last := ""
	if version == "" {
		last, _, _ = RunCommand("git", "describe", "--abbrev=0", "--tags")
		last = strings.ReplaceAll(last, "\n", "")
		version = func() string {
			subVersions := strings.Split(last, ".")
			patchVersion, _ := strconv.Atoi(subVersions[2])
			patchVersion++
			return fmt.Sprintf("%s.%s.%d", subVersions[0], subVersions[1], patchVersion)
		}()
	}

	RunCommand("git", "tag", version)
	RunCommand("git", "push", "--tags")
}
