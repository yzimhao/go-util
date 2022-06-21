package pack

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func Build(main string, dist string) {
	fmt.Println("build")

	pack := "github.com/yzimhao/utilgo/pack"
	version := "v0.0.2"
	buildTime := time.Now().Format(time.RFC3339)
	commitSha1 := getCommitSha1()
	goVersion := fmt.Sprintf("go version: %s", runtime.Version())

	fmt.Println(version, pack, buildTime, commitSha1)
	cmds := []string{
		"build",
		"-ldflags", fmt.Sprintf(`"-X %s.VERSION=%s -X %s.BUILD=%s -X %s.COMMIT_SHA1=%s -X %s.GO_VERSION="%s""`, pack, version, pack, buildTime, pack, commitSha1, pack, goVersion),
		"-o", dist,
		main,
	}
	fmt.Println(cmds)
	cmd := exec.Command("go", cmds...)
	out, err := cmd.Output()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(out))
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
