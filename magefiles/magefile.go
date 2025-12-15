package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

func isInPath(dir string) bool {
	pathEnv := os.Getenv("PATH")
	pathSeparator := ":"
	if runtime.GOOS == "windows" {
		pathSeparator = ";"
	}

	paths := strings.Split(pathEnv, pathSeparator)
	dir = filepath.Clean(dir)

	for _, p := range paths {
		cleanPath := filepath.Clean(p)
		if cleanPath == dir {
			return true
		}
	}
	return false
}

func Build() error {
	fmt.Println("Building ruv...")

	fmt.Println("Running go vet...")
	if err := sh.Run("go", "vet", "./..."); err != nil {
		return fmt.Errorf("go vet failed: %w", err)
	}

	if err := os.MkdirAll("bin", 0755); err != nil {
		return err
	}
	binary := "bin/ruv"
	if runtime.GOOS == "windows" {
		binary = "bin/ruv.exe"
	}

	// Build with optimization flags to reduce binary size:
	// -ldflags="-s -w" removes symbol table and debug info
	// -trimpath removes file system paths from binary
	return sh.Run("go", "build",
		"-ldflags=-s -w",
		"-trimpath",
		"-o", binary,
		".")
}

func getInstallDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}

	bioDir := filepath.Join(homeDir, ".bio", "bin")
	if info, err := os.Stat(bioDir); err == nil && info.IsDir() {
		return bioDir, nil
	}

	var candidateDir string
	switch runtime.GOOS {
	case "linux":
		candidateDir = filepath.Join(homeDir, ".local", "bin")
	case "windows":
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			localAppData = homeDir + "\\AppData\\Local"
		}
		candidateDir = localAppData + "\\Microsoft\\WindowsApps"
	case "darwin":
		return "", fmt.Errorf("on macOS, please create ~/.bio/bin first, or use sudo to install to /usr/local/bin")
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	// Only check PATH for platform-specific defaults
	if !isInPath(candidateDir) {
		return "", fmt.Errorf("installation directory %s is not in PATH - please create ~/.bio/bin and add it to your PATH, or add %s to your PATH", candidateDir, candidateDir)
	}

	return candidateDir, nil
}

func Install() error {
	fmt.Println("Installing ruv...")

	installDir, err := getInstallDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create install directory: %w", err)
	}

	mg.Deps(Build)

	binary := "ruv"
	if runtime.GOOS == "windows" {
		binary = "ruv.exe"
	}

	src := filepath.Join("bin", binary)
	dst := filepath.Join(installDir, binary)

	if err := sh.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to copy binary: %w", err)
	}

	if runtime.GOOS != "windows" {
		if err := os.Chmod(dst, 0755); err != nil {
			return fmt.Errorf("failed to make executable: %w", err)
		}
	}

	fmt.Printf("✓ Installed to %s\n", dst)
	return nil
}

func Clean() error {
	fmt.Println("Cleaning...")
	return sh.Rm("bin")
}

func Vet() error {
	fmt.Println("Running go vet...")
	return sh.Run("go", "vet", "./...")
}

func Run() error {
	mg.Deps(Build)
	args := os.Args[2:] // Get args after "mage run"
	binary := "bin/ruv"
	if runtime.GOOS == "windows" {
		binary = "bin/ruv.exe"
	}
	return sh.Run(binary, args...)
}
