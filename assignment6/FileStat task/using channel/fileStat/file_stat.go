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
func totalNoFiles(dict map[string]int) int {
	total := 0

	for _, v := range dict {
		total += v
	}
	return total
}

// path - path of dir
// dict - map with key-file extension, value- no of files
// wg - used to wait for go-routines until all are done
// ch - send file extension as message
// function check for each path if it represnts a file or dir
// if path is of file, just pass it through the channel
// if path is dir, then start new go-routine for it's child
// do above steps recursivly until each dir/file is processed
func findStats(path string, dict map[string]int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		childPath := path + "/" + f.Name()
		if f.IsDir() {
			wg.Add(1)
			go findStats(childPath, dict, wg, ch)
		} else {
			fileExtension := filepath.Ext(childPath)
			// adding hidden as text for hidden files
			if fileExtension == "" {
				fileExtension = "hidden"
			}
			//m.Lock()
			//dict[fileExtension]++
			//m.Unlock()
			ch <- fileExtension
		}
	}
}

// read dir path from console as input
func ReadDirPath() string {
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

// populate map[file extension]no_of_files
func LoadDirStat(path string) {
	// map with file extension as key and count as value
	dict := make(map[string]int)

	wg := &sync.WaitGroup{} // this will help this func to wait until every go-routine finished it's job
	ch := make(chan string) // this will help us to pass file extension of each file

	// 1 for findStat and 1 for wait function
	wg.Add(2)
	go findStats(path, dict, wg, ch)

	// because we don't know total no of go-routines
	// this will close the close the as soon as all go-routine finish their work
	go func() {
		wg.Done()
		wg.Wait()
		close(ch)
	}()

	func() {
		for msg := range ch { // putting very file extension recived from channel to this map
			dict[msg]++
		}
	}()

	printStat(dict)
}

// print file extension with no of files of each type
// print total no of files
func printStat(dict map[string]int) {
	fmt.Printf("%30s%15s\n", "Extension", "No of files")
	for key, value := range dict {
		fmt.Printf("%30s%15d\n", key, value)
	}
	fmt.Println("Total no of files: ", totalNoFiles(dict))
}
