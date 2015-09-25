package main

import (
    "fmt"
    "github.com/go-yaml/yaml"
    "io/ioutil"
	"path/filepath"
	"log"
	"flag"
)

type PaintList struct {
	Description string
    Gwpaint map[string][]string
	Papaint map[string][]string
}

func handleError(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
	return
}

func main() {
	var filename string
	var plist PaintList
	
	flag.StringVar(&filename, "file", "", "a YAML config file")
	
	flag.Parse()
	
	if len(filename) == 0 {
		log.Fatalln("[ERROR] - please read usage through -h or --help option.")
	}
	
    file, err := filepath.Abs(filename)
	handleError(err)
	
	source, err := ioutil.ReadFile(file)
	handleError(err)	
	
	err = yaml.Unmarshal(source, &plist)
	handleError(err)
	
    fmt.Printf("Description: %#v\n", plist.Description)
	
	for key, value := range plist.Gwpaint {
    fmt.Println("Base:", key, "Layer:", value)
	}
	return
}