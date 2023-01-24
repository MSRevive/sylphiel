package main

import (
	"time"
	"fmt"
	"os"

	"github.com/msrevive/sylphiel/cmd"
	"github.com/msrevive/sylphiel/internal/system"
)

var spMsg string = `
Sylphiel (Syl-phe-iel).

Copyright © %d, Team MSRebirth
Version: %s
Website: https://msrebirth.net/
License: GPL-3.0 https://github.com/MSRevive/sylphiel/blob/main/LICENSE %s
`

func main() {
	fmt.Printf(spMsg, time.Now().Year(), system.Version, "\n\n")

	if err := cmd.Run(os.Args); err != nil {
		panic(err)
	}
}