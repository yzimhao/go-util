package build

import (
	"fmt"
	"os/exec"
	"time"
)

func Build(main string) {
	fmt.Println("build")

	version := "v0.0.2"
	goUtil := "github.com/yzimhao/go-util"
	buildTime := time.Now().Format(time.RFC3339)
	commitSha1 := getCommitSha1()

	fmt.Println(version, goUtil, buildTime, commitSha1)
	cmd := exec.Command("go", "build", "-o", "dist/build", main)
	cmd.Output()
}

func getCommitSha1() string {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}
