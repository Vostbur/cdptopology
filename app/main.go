package main

import (
	"cdptopology/app/tools"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := tools.CheckArgs(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(0)
	}

	dir := os.Args[1]
	files, err := tools.GetSourceFiles(dir)
	if err != nil {
		log.Fatal(fmt.Errorf("Error reading files from dir %s. %v", dir, err))
	}

	devs, devMap := tools.GetDevsInfo(tools.ReadSourceFiles(files))

	err = tools.BuildJson(devs, devMap)
	if err != nil {
		log.Fatal(fmt.Errorf("Error building JSON. %v", err))
	}
}
