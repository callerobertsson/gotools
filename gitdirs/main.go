// Package main implement the gitdirs tool for finding git directories in a file structure.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

const (
	gitDirName = ".git"
)

var (
	basePath       string = "."
	onlyGitDirs    bool
	onlyNonGitDirs bool
	recursive      bool
	silent         bool
)

func init() {
	flag.BoolVar(&onlyGitDirs, "g", false, "Only show git directories")
	flag.BoolVar(&onlyNonGitDirs, "G", false, "Only show non-git directories")
	flag.BoolVar(&recursive, "r", false, "Search git dirs in subdirectories")
	flag.BoolVar(&silent, "s", false, "Only print paths (and prefix if not -g or -G)")

	flag.Parse()

	if onlyGitDirs && onlyNonGitDirs {
		fmt.Println("Error: The -g and -G flags are mutually exclusive")
		os.Exit(1)
	}

	if len(flag.Args()) > 0 {
		basePath = flag.Arg(0)
	}
}

func main() {
	gitDirs(basePath)
}

func gitDirs(d string) {
	fis, err := ioutil.ReadDir(d)
	if err != nil {
		fmt.Printf("Error: Could not read directory %s\n", d)
		return
	}

	for _, fi := range fis {
		if !fi.IsDir() {
			continue
		}

		gitPrefix := ""
		nonGitPrefix := ""
		if !silent || !onlyGitDirs && !onlyNonGitDirs {
			gitPrefix = "GIT: "
			nonGitPrefix = "NON: "
		}

		path := d + string(os.PathSeparator) + fi.Name()
		isGit := isGitDir(path)

		switch {
		case isGit && !onlyNonGitDirs:
			branch := ""
			status := ""
			if !silent {
				branch, status = gitStatus(path)
				branch = " [" + branch + "]"
				status = " " + status + " "
			}
			fmt.Printf("%s%s%s%s\n", gitPrefix, status, path, branch)
		case !isGit && !onlyGitDirs:
			fmt.Printf("%s%s\n", nonGitPrefix, path)
		}

		if !isGit && recursive {
			gitDirs(path)
		}
	}
}

func isGitDir(path string) bool {

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error: Directory %s not readable\n", path)
		return false
	}

	for _, fi := range fis {
		if fi.IsDir() && fi.Name() == gitDirName {
			return true
		}
	}

	return false
}

func gitStatus(p string) (branch, status string) {

	cmd := exec.Command("git", "status")
	cmd.Dir = p
	bs, err := cmd.CombinedOutput()
	if err != nil {
		return branch, "Could not do git status"
	}
	out := string(bs)

	branchRegex := regexp.MustCompile(`^On branch ([^\s]+)`)
	branch = branchRegex.FindAllStringSubmatch(out, -1)[0][1]

	cleanRegex := regexp.MustCompile(`nothing to commit`)
	status = "Changes"
	if cleanRegex.MatchString(out) {
		status = "Clean  "
	}

	return branch, status
}
