// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-FileCopyrightText: 2024 kurth4cker <kurth4cker@gmail.com>

package main

func main() {
	var project Project
	project.Initialize()
	project.Scan()

	project.Build()
}
