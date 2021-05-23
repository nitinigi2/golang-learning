package main

import (
	"fmt"
	"time"

	"github.com/nitinigi2/assignment6/fileStat"
)

//directoryPath = "C://Program Files/Go"
func main() {

	dirPath := fileStat.ReadDirPath()
	// logging start and end time
	start := time.Now()

	fileStat.LoadDirStat(dirPath)

	elapsed := time.Since(start)
	fmt.Printf("process took %s", elapsed)
}
