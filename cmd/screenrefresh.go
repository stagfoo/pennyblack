package main

import "os/exec"

func ScreenRefresh() {
	cmd := exec.Command("import", "-window", "root",
		"-display", ":99",
		"screendump/current.png")
	cmd.Run()
}
