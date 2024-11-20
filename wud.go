// Author: Madhanmaaz
// Website: https://madhanmaaz.netlify.app/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func disableWindowsUpdate() error {
	// Check if the program has administrative privileges
	if !isAdmin() {
		return fmt.Errorf("this program needs to be run as an administrator")
	}

	// Command to disable the Windows Update service
	cmd := exec.Command("cmd", "/C", `sc config wuauserv start= disabled & net stop wuauserv /y`)

	// Hide the command prompt window
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	// run the command
	cmd.Run()

	fmt.Println("Windows Update service has been disabled successfully.")
	return nil
}

func isAdmin() bool {
	// Check if the process is running with administrator privileges
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func main() {
	delay := 1 * time.Minute // Start with a 1-minute delay

	for {
		// Call the function
		err := disableWindowsUpdate()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		// Wait for the delay
		fmt.Printf("Waiting for %v before the next call...\n", delay)
		time.Sleep(delay)
	}
}

// run command: go build -ldflags -H=windowsgui -o wud.exe wud.go
