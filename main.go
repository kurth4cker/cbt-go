// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2024 kurth4cker <kurth4cker@gmail.com>

package main

import (
	"log"
	"path/filepath"
)

type Compiler string

func (c *Compiler) SetCommand(cmd string) {
	*c = Compiler(cmd)
}

func (c *Compiler) Compile(src string) {
}

func (c *Compiler) Link(objects []string, output string) {
}

func (c *Compiler) CompileAndLink(sources []string, output string) {
}

// top struct which holds all project data
type Project struct {
	Name string

	// Where binary will be put
	BuildDir string

	// C Compiler [cc]
	CC Compiler

	RootDir string // Project root
	Executables []string
	Libraries []string
	IncludeDirs []string
}

func (p *Project) Build() {
}

func (p *Project) AddExecutable(exec string) {
}

func (p *Project) AddSource(src string) {
}

func (p *Project) AddLibrary(lib string) {
}

func (p *Project) AddIncludeDir(dir string) {
}

func (p *Project) SetBuildDir(dir string) {
}

func (p *Project) SetCC(cc *Compiler) {
}

func (p *Project) SetDefaults() {
}

func (p *Project) SetName(name string) {
}

func (p *Project) SetRootDir(dir string) {
}

var DefaultProject = Project {
	BuildDir: ".",
}

func main() {
	var project Project
	project.SetDefaults()

	// TODO: maybe these should be in project.SetDefaults
	sourceFiles, _ := filepath.Glob("*.c")
	for _, source := range sourceFiles {
		project.AddSource(source)
	}

	// TODO: also these: project.SetDefaults?
	buildDir, err :=  filepath.Abs(project.BuildDir)
	if err != nil {
		log.Fatalln(err)
	}
	project.SetBuildDir(buildDir)

	project.Build()
}
