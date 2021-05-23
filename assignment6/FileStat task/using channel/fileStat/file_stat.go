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

func totalNoFiles(dict map[string]int) int {
	total := 0

	for _, v := range dict {
		total += v
	}
	return total
}

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

func LoadDirStat(path string) {
	// map with file extension as key and count as value
	dict := make(map[string]int)

	wg := &sync.WaitGroup{}
	ch := make(chan string)

	// 1 for findStat and 1 for wait function
	wg.Add(2)
	go findStats(path, dict, wg, ch)

	go func() {
		wg.Done()
		wg.Wait()
		close(ch)
	}()

	func() {
		for msg := range ch {
			dict[msg]++
		}
	}()

	printStat(dict)
}

func printStat(dict map[string]int) {
	fmt.Printf("%30s%15s\n", "Extension", "No of files")
	for key, value := range dict {
		fmt.Printf("%30s%15d\n", key, value)
	}
	fmt.Println("Total no of files: ", totalNoFiles(dict))
}
