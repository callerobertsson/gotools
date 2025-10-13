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
	pullAll        bool
	switchBranch   string
	recursive      bool
	silent         bool
)

// Terminal text and color
const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
	Strike    = "\033[9m"
	Italic    = "\033[3m"

	CRed    = "\033[31m"
	CGreen  = "\033[32m"
	CYellow = "\033[33m"
	CBlue   = "\033[34m"
	CPurple = "\033[35m"
	CCyan   = "\033[36m"
	CWhite  = "\033[37m"
)

func init() {
	flag.BoolVar(&onlyGitDirs, "g", false, "Only show git directories")
	flag.BoolVar(&onlyNonGitDirs, "G", false, "Only show non-git directories")
	flag.StringVar(&switchBranch, "b", "", "Try to switch to branch in all repos")
	flag.BoolVar(&pullAll, "p", false, "Execute pull in all git directories")
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
		fmt.Printf(CRed+"Error: Could not read directory %s\n"+Reset, d)
		return
	}

	for _, fi := range fis {
		if !fi.IsDir() {
			continue
		}

		gitPrefix := ""
		nonGitPrefix := ""
		if !silent || !onlyGitDirs && !onlyNonGitDirs {
			gitPrefix = CYellow + "GIT: " + Reset
			nonGitPrefix = CWhite + "NON: " + Reset
		}

		path := d + string(os.PathSeparator) + fi.Name()
		isGit := isGitDir(path)

		switch {
		case isGit && !onlyNonGitDirs:
			if switchBranch != "" {
				output, err := gitCheckoutBranch(path, switchBranch)
				if err != nil {
					fmt.Printf(CRed+"Set Branch Error: %s\n"+Reset, err)
				}
				fmt.Printf("%s%s\n", "  ", CWhite+output+Reset)
			}

			if pullAll {
				output, err := gitPull(path)
				if err != nil {
					fmt.Printf(CRed+"Error: %v\n"+Reset, err)
				}
				fmt.Printf("%s%s\n", "  ", CWhite+output+Reset)
			}
			branch := ""
			status := ""
			if !silent {
				branch, status = gitStatus(path)
				branch = " [" + branch + "]"
				status = " " + status + " "
			}
			fmt.Printf("%s%s%s%s\n", gitPrefix, status, CPurple+path+Reset, CWhite+branch+Reset)
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

func gitCheckoutBranch(p, b string) (string, error) {

	cmd := exec.Command("git", "checkout", b)
	cmd.Dir = p

	bs, _ := cmd.CombinedOutput()

	return string(bs), nil
}

func gitPull(p string) (string, error) {

	cmd := exec.Command("git", "pull")
	cmd.Dir = p

	bs, _ := cmd.CombinedOutput()

	return string(bs), nil
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
	status = CRed + "Changes" + Reset
	if cleanRegex.MatchString(out) {
		status = CGreen + "Clean  " + Reset
	}

	return branch, status
}
