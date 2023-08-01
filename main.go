package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	urlFlag := flag.String("url", "", "URL of the repository to clone")
	flag.Parse()

	// Check if the URL flag is provided
	if *urlFlag == "" {
		fmt.Println("Please provide a URL using -url flag.")
		return
	}

	// Create a temporary directory to clone the repository
	tempDir, err := os.MkdirTemp("", "git-clone-")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		return
	}

	fmt.Println("temp dir: ", tempDir)

	cmd := exec.Command("git", "clone", *urlFlag, tempDir)
	output, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error cloning repository: %v\nOutput: %s\n", cmdErr, output)
		return
	}

	fmt.Println("Repository cloned successfully!")
	fmt.Println("Cloned repository path:", tempDir)

	coverProfilePath := filepath.Join(tempDir, "cover.out")
	os.Chdir(tempDir)
	testCmd := exec.Command("go", "test", "-coverprofile", coverProfilePath, "./...")
	testOutput, testCmdErr := testCmd.CombinedOutput()
	if testCmdErr != nil {
		fmt.Printf("Error running go test: %v\nOutput: %s\n", testCmdErr, testOutput)
		return
	}

	fmt.Println("Tests executed successfully!")
	fmt.Println("Test coverage profile generated at:", filepath.Join(tempDir, "cover.out"))

	// Execute 'go tool cover' command to generate HTML coverage report
	coverageHTMLPath := filepath.Join(tempDir, "coverage.html")
	coverCmd := exec.Command("go", "tool", "cover", "-html", coverProfilePath, "-o", coverageHTMLPath)
	coverOutput, coverCmdErr := coverCmd.CombinedOutput()
	if coverCmdErr != nil {
		fmt.Printf("Error generating coverage report: %v\nOutput: %s\n", coverCmdErr, coverOutput)
		return
	}

	fmt.Println("HTML coverage report generated at:", coverageHTMLPath)

	// Open the HTML coverage report in the default browser
	if err := openBrowser(coverageHTMLPath); err != nil {
		fmt.Println("Error opening coverage report in browser:", err)
	}
}

func openBrowser(htmlFilePath string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("open", htmlFilePath).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	return err
}
