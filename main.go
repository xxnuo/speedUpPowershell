package main

import (
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("powershell", "-Command", "Write-Host", "SpeedUp")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	if cmd.Run() == nil {
		println("ok")
	} else {
		println("err")
	}
}

// C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp
