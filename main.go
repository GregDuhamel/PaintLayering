package main

import (
    "gopkg.in/yaml.v2"
	"gopkg.in/mgo.v2"
    "io/ioutil"
	"path/filepath"
	"log"
	"flag"
	"fmt"
	"time"
)

type Config struct {
	Descritption string
	Url map[string]string
}

func handleError(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
	return
}

func buildURL(conf *Config) string {
	if (len(conf.Url["login"]) == 0) {
		conf.Url["login"] = "Guest"
	}
	if (len(conf.Url["password"]) == 0) {
		conf.Url["password"] = "Guest"
	}
	if (len(conf.Url["hostname"]) == 0) {
		conf.Url["hostname"] = "localhost"
	}
	if (len(conf.Url["port"]) == 0) {
		conf.Url["port"] = "27017"
	}
	if (len(conf.Url["database"]) == 0) {
		conf.Url["database"] = "PaintLayering"
	}
	if (len(conf.Url["options"]) == 0) {
		conf.Url["options"] = ""
	}
	conn := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?%s", conf.Url["login"], conf.Url["password"], conf.Url["hostname"], conf.Url["port"], conf.Url["database"], conf.Url["options"])
	
	return conn
}

func main() {
	var filename string
	var conf Config
	
	flag.StringVar(&filename, "conf", "", "a YAML config file")
	
	flag.Parse()
	
	if len(filename) == 0 {
		log.Fatalln("[ERROR] - please read usage through -h or --help option.")
	}
	
    file, err := filepath.Abs(filename)
	handleError(err)
	
	source, err := ioutil.ReadFile(file)
	handleError(err)
	
	err = yaml.Unmarshal(source, &conf)
	handleError(err)
	
	url := buildURL(&conf)
	
	fmt.Println(url)
	
	session, err := mgo.DialWithTimeout(url, 15 * time.Second)
	handleError(err)

	fmt.Println(session)

	return
}