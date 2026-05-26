package main

import (
	"github.com/dylannorthrup/go-internetarchive/cmd"
	iaInt "github.com/dylannorthrup/go-internetarchive/internal"
)

// Version takes in version string from build_all.sh
var Version = "vMAIN"

func main() {
	iaInt.DieOnError(cmd.Root.Execute())
}
