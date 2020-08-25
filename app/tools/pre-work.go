package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

//test
func CheckArgs() (err error) {
	if len(os.Args) != 2 {
		err = errors.New(`
		Usage: cdptopology.exe DIR
		
			   DIR - a folder with files that is the result of executing
					 the 'sh cdp neighboors' command on each device`)
	}
	return
}

// test
func GetSourceFiles(dir string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			continue
		}
		files = append(files, path.Join(dir, file.Name()))
	}
	return files, nil
}

// test
func ReadSourceFiles(files []string) <-chan string {
	wg := &sync.WaitGroup{}
	rawDataChan := make(chan string, len(files))
	for _, f := range files {
		wg.Add(1)
		go func(f string, wg *sync.WaitGroup) {
			defer wg.Done()
			rawData, err := ioutil.ReadFile(f)
			if err != nil {
				fmt.Printf("error reading file %s\n", f)
				rawDataChan <- ""
				return
			}
			rawDataChan <- string(rawData)
		}(f, wg)
	}
	wg.Wait()
	close(rawDataChan)
	return rawDataChan
}
