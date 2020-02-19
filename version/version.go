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
	Version string
	Commit  string
	Date    string
)

// PrintVersion returns the version for the command version and --version flag
func PrintVersion(basename string) {
	fmt.Printf("%s version: %s, build: %s, date: %s\n", basename, Version, Commit, Date)
	os.Exit(0)
}
