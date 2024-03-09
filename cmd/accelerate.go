package main

import (
	// "flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/MichaelCduBois/accelerate/internal/generator"
	"github.com/MichaelCduBois/accelerate/internal/helpers"
)

func main() {
	// Check for minimum number of arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: accelerate <module> <directory>\n\nExample:\n\taccelerate github.com/MichaelCduBois/rickroll rickroll")
		return
	}
	// Create project directory
	cmd := exec.Command("mkdir", os.Args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return
	}
	// Create git repository
	cmd = exec.Command("git", "init")
	cmd.Dir = os.Args[2]
	if err := cmd.Run(); err != nil {
		return
	}
	// Change to project directory
	cmd = exec.Command("cd", os.Args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Unable to change to directory: ", err)
		return
	}
	// Initialize go module
	cmd = exec.Command("go", "mod", "init", os.Args[1])
	cmd.Dir = os.Args[2]
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Unable to initialize go module: ", err)
		return
	}
	// Get Installed Go Version
	goVersion, err := helpers.GetGoVersion()
	if err != nil {
		fmt.Println("Unable to get go version: ", err)
		return
	}
	// Generate Docker Config Files
	generator.DockerConfigFiles(string(goVersion), os.Args[2])
}
