// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2024 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"log"
	"os/exec"
	"path/filepath"
	"slices"
)

// top struct which holds all project data
type Project struct {
	Name string
	CC string
	Sources []string
}

// Build the project
func (p *Project) Build() {
	args := []string{"-o", p.Name}
	args = slices.Concat(args, p.Sources)

	cmd := exec.Command(p.CC, args...)
	cmd.Run()
	log.Println(cmd)
}

// Set default values
func (p *Project) Initialize() {
	p.CC = "cc"
}

// Scan and set corresponding values. Executables, project name, libraries etc.
// Initialize project structure
func (p *Project) Scan() {
	absolute_path, err := filepath.Abs(".")
	if err != nil {
		log.Fatalln(err)
	}
	p.Name = filepath.Base(absolute_path)

	p.Sources, _ = filepath.Glob("./*.c")
}
