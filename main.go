/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/zachdehooge/GO-SAST-Scanner/cmd"
)

func isGoSecInstalled() bool {
	cmd := exec.Command("which", "gosec")
	err := cmd.Run()
	return err == nil
}

// installGoSec installs GoSec using the Go install command
func installGoSec() error {
	// Command to install gosec
	cmd := exec.Command("go", "install", "github.com/securego/gosec/v2@latest")
	cmd.Stdout = exec.Command("tee").Stdout // Redirect output to stdout
	cmd.Stderr = exec.Command("tee").Stderr // Redirect error to stderr

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error installing GoSec: %v, output: %s", err, output)
	}
	return nil
}

func goSecDependency() {
	if !isGoSecInstalled() {
		fmt.Println("GoSec is not installed. Installing...\n")

		// Install gosec
		if err := installGoSec(); err != nil {
			log.Fatalf("Failed to install GoSec: %s", err)
		}

		fmt.Println("GoSec has been installed successfully.")
	}
}

func isSnykInstalled() bool {
	cmd := exec.Command("which", "snyk")
	err := cmd.Run()
	return err == nil
}

// installGoSec installs GoSec using the Go install command
func installSnyk() error {
	// Command to install gosec
	cmd := exec.Command("npm", "install", "snyk", "-g")
	cmd.Stdout = exec.Command("tee").Stdout // Redirect output to stdout
	cmd.Stderr = exec.Command("tee").Stderr // Redirect error to stderr

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error installing Snyk: %v, output: %s", err, output)
	}
	return nil
}

func snykDependency() {
	if !isSnykInstalled() {
		fmt.Println("Snyk is not installed. Installing...\n")

		// Install gosec
		if err := installSnyk(); err != nil {
			log.Fatalf("Failed to install Snyk: %s", err)
		}

		fmt.Println("Snyk has been installed successfully.")
	}
}

func main() {
	goSecDependency()
	snykDependency()
	cmd.Execute()
}
