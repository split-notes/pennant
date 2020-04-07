package bash

import "fmt"

// Base
func GitStash() string     { return "git stash" }
func GitResetHard() string { return "git reset HEAD --hard" }
func GitClean() string     { return "git clean -fd" }
func GitOnboard() string   { return "git submodule update --init" }
func GitBranch() string    { return "git checkout master" }

// Recursive
func GitFetchRecurse() string     { return "git fetch --recurse-submodules" }
func GitPullRecurse() string      { return "git pull --recurse-submodules" }
func GitStashRecurse() string     { return fmt.Sprintf("git submodule foreach '%s'", GitStash()) }
func GitResetHardRecurse() string { return fmt.Sprintf("git submodule foreach '%s'", GitResetHard()) }
func GitCleanRecurse() string     { return fmt.Sprintf("git submodule foreach '%s'", GitClean()) }
func GitBranchRecurse() string    { return fmt.Sprintf("git submodule foreach '%s'", GitBranch()) }
