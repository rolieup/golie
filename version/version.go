/*
Copyright © 2020 Rolie and Golie Contributors. See LICENSE for license.
*/

package version

import (
	"fmt"
	"os"
)

// Version is the version of the build.
// Build details
var (
	// For testing purposes, let's hard code Version and Date.
	// TBD: Automated method for updating this file during release/tagging
	Version = "v0.2.1"
	// Short month notation
	Date = "2020-Oct-15"
	Commit string
)

// PrintVersion returns the version for the command version and --version flag
func PrintVersion(basename string) {
	fmt.Printf("%s version: %s, build: %s, date: %s\n", basename, Version, Commit, Date)
	os.Exit(0)
}
