package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	fmt.Println(path)
}
