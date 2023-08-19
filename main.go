package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "", "URL of the repository to clone")
	pathFlag := flag.String("path", "", "Path of the directory")
	flag.Parse()

	if *urlFlag == "" && *pathFlag == "" {
		fmt.Println("Path or Url are not provided")
		return
	}

	if *urlFlag != "" && *pathFlag != "" {
		fmt.Println("Cannot use both Path and Url")
		return
	}

	tempDir := *pathFlag
	if *urlFlag != "" {
		var err error
		tempDir, err = os.MkdirTemp("", "git-clone-")
		if err != nil {
			fmt.Println("Error creating temporary directory:", err)
			return
		}
		fmt.Println("temp dir:", tempDir)

		if cloneErr := cloneToPath(*urlFlag, tempDir); cloneErr != nil {
			fmt.Println("Error cloning the repo: ", cloneErr)
			return
		}
	}

	coverProfilePath := filepath.Join(tempDir, "cover.out")
	os.Chdir(tempDir)
	testCmd := exec.Command("go", "test", "-coverprofile", coverProfilePath, "./...")
	testOutput, testCmdErr := testCmd.CombinedOutput()
	if strings.Contains(string(testOutput), "no test files") {
		fmt.Println("Tests don't exist")
		return
	}
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

func cloneToPath(url, path string) error {
	cmd := exec.Command("git", "clone", url, path)
	output, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		fmt.Printf("Error cloning repository: %v\nOutput: %s\n", cmdErr, output)
		return cmdErr
	}

	fmt.Println("Repository cloned successfully!")
	fmt.Println("Cloned repository path:", path)

	return nil
}
