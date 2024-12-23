package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

// ProcessInfo holds information about a process, such as its PID and name
type ProcessInfo struct {
	PID  string
	Name string
}

// findAndKillProcess searches for processes by name and kills them
func findAndKillProcess(processName string) error {
	var cmd *exec.Cmd
	var output []byte
	var err error

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("ps", "-A", "-o", "pid,comm")
		output, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("failed to get process list: %v", err)
		}
	case "windows":
		cmd = exec.Command("tasklist")
		output, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("failed to get process list: %v", err)
		}
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// Parse the output and find the process by name
	processes, err := parseProcesses(output, processName)
	if err != nil {
		return err
	}

	if len(processes) == 0 {
		return fmt.Errorf("no process found with the name: %s", processName)
	}

	for _, process := range processes {
		err = killProcess(process.PID)
		if err != nil {
			log.Printf("Error killing process %s (PID %s): %v", process.Name, process.PID, err)
		}
	}

	return nil
}

// parseProcesses parses the output of the system process listing commands and returns a list of matching processes
func parseProcesses(output []byte, processName string) ([]ProcessInfo, error) {
	var processes []ProcessInfo
	lines := strings.Split(string(output), "\n")

	// Skip header lines if present (on some systems, ps command has headers)
	startLine := 0
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		if len(lines) > 1 && strings.HasPrefix(lines[0], "PID") {
			startLine = 1 // Skip header
		}
	}

	// Iterate through the process list and match the process name
	for _, line := range lines[startLine:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		columns := strings.Fields(line)
		if len(columns) < 2 {
			continue // Skip invalid lines
		}

		pid := columns[0]
		name := columns[1]

		// Match by partial or exact name
		if strings.Contains(name, processName) {
			processes = append(processes, ProcessInfo{PID: pid, Name: name})
		}
	}

	if len(processes) == 0 {
		return nil, fmt.Errorf("no processes found matching the name: %s", processName)
	}

	return processes, nil
}

// killProcess sends a kill signal to the process by PID
func killProcess(pid string) error {
	var cmd *exec.Cmd

	// Based on the OS, use the appropriate kill command
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("kill", pid)
	case "windows":
		cmd = exec.Command("taskkill", "/PID", pid)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to kill process %s: %v", pid, err)
	}

	return nil
}

// validateInput checks if the user has provided a valid process name
func validateInput(args []string) error {
	if len(args) < 2 {
		return errors.New("you must specify a process name")
	}
	return nil
}
