package tools

import (
	"fmt"
	"regexp"
	"sync"
)

type Neighboors struct {
	Hostname        string
	LocalInterface  string
	RemoteInterface string
}

type Device struct {
	Hostname   string
	Neighboors []Neighboors
}

// test
func GetDevsInfo(c <-chan string) ([]Device, map[string]int) {
	dc := parser(c)
	var devs []Device
	devMap := make(map[string]int, len(dc))
	var ind int
	for d := range dc {
		devs = append(devs, d)
		devMap[d.Hostname] = ind
		ind++
	}
	return devs, devMap
}

// test
func parser(contentChan <-chan string) <-chan Device {
	wg := &sync.WaitGroup{}
	devChan := make(chan Device, len(contentChan))
	for content := range contentChan {
		wg.Add(1)

		go func(r string, c chan Device, wg *sync.WaitGroup) {
			defer wg.Done()
			var tmp Device
			var err error
			tmp.Hostname, err = parseHostname(r)
			if err != nil {
				fmt.Println(err.Error())
				c <- Device{}
				return
			}
			var tmpStrSlice [][]string
			tmpStrSlice, err = parseShowCDP(r)
			if err != nil {
				fmt.Println(err.Error())
				c <- Device{}
				return
			}
			for _, i := range tmpStrSlice {
				if i[2] != "Fas 0/0" { //HOOK for switch
					tmp.Neighboors = append(tmp.Neighboors, Neighboors{i[1], i[2], i[3]})
				}
			}
			c <- tmp
		}(content, devChan, wg)
	}
	wg.Wait()
	close(devChan)
	return devChan
}

// test
func parseHostname(raw string) (string, error) {
	re, err := regexp.Compile(`(\w+?)[>|#]`)
	if err != nil {
		return "", err
	}
	match := re.FindStringSubmatch(raw)
	if len(match) != 2 {
		return "", nil
	}
	return match[1], nil
}

// test
func parseShowCDP(show string) ([][]string, error) {
	re, err := regexp.Compile(`(?im)(^\w+)\s+(\w+\s\d+/\d+) .+? (\w+\s\d+/\d+)`)
	if err != nil {
		return nil, err
	}
	result := re.FindAllStringSubmatch(show, -1)
	return result, nil
}
