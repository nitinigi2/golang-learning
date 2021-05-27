package fileStat

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// return total no of files in the input directory
// populate map[file extension]no_of_files
func GetDirStats(directoryPath string) {
	// map with file extension as key and count as value
	dict := make(map[string]int)

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	wg.Add(1)
	go findStats(directoryPath, dict, wg, m)
	wg.Wait()

	printStats(dict)
}

// print file extension with no of files of each type
// print total no of files
func printStats(dict map[string]int) {
	fmt.Printf("%30s%15s\n", "Extension", "No of files")
	for key, value := range dict {
		fmt.Printf("%30s%15d\n", key, value)
	}
	fmt.Println("Total no of files: ", totalNoFiles(dict))
}

// return total no of files in the input directory
func totalNoFiles(dict map[string]int) int {
	total := 0

	for _, v := range dict {
		total += v
	}
	return total
}

// read dir path from console as input
func ReadDirectoryPath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter directory path: ")
	directoryPath, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// On Windows, the line is terminated with "\r\n", so remove that.
	directoryPath = strings.TrimSuffix(directoryPath, "\r\n")

	return directoryPath
}

// path - path of dir
// dict - map with key-file extension, value- no of files
// wg - used to wait for go-routines until all are done
// m - mutex, used for locking map
// function check for each path if it represnts a file or dir
// if path is of file, just pass it through the channel
// if path is dir, then start new go-routine for it's child
// do above steps recursivly until each dir/file is processed
func findStats(path string, dict map[string]int, wg *sync.WaitGroup, m *sync.Mutex) {

	defer wg.Done()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		childPath := path + "/" + f.Name()
		if f.IsDir() {
			wg.Add(1)
			go findStats(childPath, dict, wg, m)
		} else {
			fileExtension := filepath.Ext(childPath)
			// adding hidden as text for hidden files
			if fileExtension == "" {
				fileExtension = "hidden"
			}

			// apply lock as concurrent writes are not allowed
			m.Lock()
			dict[fileExtension]++
			m.Unlock()
		}
	}
}
