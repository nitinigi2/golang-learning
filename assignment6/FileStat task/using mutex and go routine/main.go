package main

import (
	"fmt"
	"time"

	"github.com/nitinigi2/assignment6/fileStat"
)

//directoryPath = "C://Program Files/Go"
func main() {

	directoryPath := fileStat.ReadDirectoryPath()

	// logging start and end time
	start := time.Now()

	fileStat.GetDirStats(directoryPath)
	elapsed := time.Since(start)
	fmt.Printf("Total time took %s", elapsed)
}
