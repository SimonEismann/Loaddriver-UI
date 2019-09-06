package shared

import (
	"os"
	"path/filepath"
)

// HomeDir is used to get the platform specific HOME directory
func HomeDir() string {
	// *nix systems
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	// windows
	return os.Getenv("USERPROFILE")
}

// Must is used to enforce a panic in case the given error is not nil
func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// GetExecDir retrieves the file directory of the running executable
func GetExecDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	Must(err)
	return dir
}

// MinInIntSlice returns the minimum value over the values inside the given int slice
func MinInIntSlice(ints []int) int {
	var min int
	for i, e := range ints {
		if i == 0 || e < min {
			min = e
		}
	}
	return min
}
