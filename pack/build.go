package pack

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const defaultFailedCode = 1

func RunCommand(name string, args ...string) (stdout string, stderr string, exitCode int) {
	log.Println("run command:", name, args)
	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		// try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	log.Printf("command result, stdout: %v, stderr: %v, exitCode: %v", stdout, stderr, exitCode)
	return
}

func Build(main string, dist string) {
	fmt.Println("build")

	pack := "github.com/yzimhao/utilgo/pack"
	version := "v0.0.3"
	buildTime := time.Now().Format(time.RFC3339)
	commitSha1 := strings.ReplaceAll(getCommitSha1(), "\n", "")
	// goVersion := fmt.Sprintf("go version: %s", runtime.Version())

	fmt.Println(version, pack, buildTime, commitSha1)
	cmds := []string{
		"gox",
		// "-ldflags", `"-s -w"`,
		"-ldflags", fmt.Sprintf(`"-X %s.VERSION=%s"`, pack, version),
		`-osarch`, "linux/amd64",
		`-osarch`, "darwin/amd64",
		"-output", fmt.Sprintf("%s_%s", dist, version),
		main,
	}

	// cmd := exec.Command(cmds[0], cmds[1:]...)

	// fmt.Println(cmd.String())

	// var out bytes.Buffer
	// var stderr bytes.Buffer

	// cmd.Stdout = &out
	// cmd.Stderr = &stderr
	// if err := cmd.Run(); err != nil {
	// 	fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	// 	panic(err)
	// }
	// fmt.Println(out.String())

	RunCommand(cmds[0], cmds[1:]...)
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
