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

func printStats(dict map[string]int) {
	fmt.Printf("%30s%15s\n", "Extension", "No of files")
	for key, value := range dict {
		fmt.Printf("%30s%15d\n", key, value)
	}
	fmt.Println("Total no of files: ", totalNoFiles(dict))
}

func totalNoFiles(dict map[string]int) int {
	total := 0

	for _, v := range dict {
		total += v
	}
	return total
}

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

/*
	if current path represnts a dir then start a new go-routine and
	call this function recursively
	if current path represnts a file then update map[file_extension]no_of_files
*/
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
			m.Lock()
			dict[fileExtension]++
			m.Unlock()
		}
	}
}
