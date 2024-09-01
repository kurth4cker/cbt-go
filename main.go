// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2024 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

// support only current dir for now
var buildDir = "."

var CC = "cc"

func main() {
	sourceFiles, _ := filepath.Glob("*.c")

	buildDir, err :=  filepath.Abs(buildDir)
	if err != nil {
		log.Fatalln(err)
	}
	outputName := filepath.Base(buildDir)

	var cmd = exec.Command(CC)
	cmd.Args = append(cmd.Args, sourceFiles...)
	cmd.Args = append(cmd.Args, "-o")
	cmd.Args = append(cmd.Args, outputName)
	cmd.Run()
	fmt.Println(cmd)
}
