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
	"strconv"
	"net/http"
	"encoding/json"
	"os"
)

type Config struct {
	Descritption string
	MongoDB map[string]string
	Url map[string]string
}

type envId struct {
	CustomLimitMessage	string `json:"customLimitMessage"`
	Skus				[]paintId `json:"skus"`
	Dangerous 			bool `json:"dangerous"`
    Hazardous 			bool `json:"hazardous"`
	LimitedQuantity 	uint8 `json:"limitedQuantity"`
	ProductId			string `json:"productId"`
}

type paintId struct {
	Id string `json:"id"`
	AvailabilityStatus string `json:"availabilityStatus"`
	Title string `json:"title"`
	Price string `json:"price"`
	ImageName string `json:"imageName"`
	ProductTitle string `json:"productTitle"`
	DisplayName string `json:"displayName"`
	Swatch string `json:"swatch"`
	Dangerous bool `json:"dangerous"`
	ProductType string `json:"productType"`
	ProductId string `json:"productId"`
}

func handleError(e error) {
	if e != nil {
		log.Fatalf("[ERROR] - %v", e)
	}
	return
}

type plError struct {
	When time.Time
	What string
}

func (e *plError) Error() string {
	return fmt.Sprintf("at %v, %s",
	e.When, e.What)
}

func buildURL(conf *Config) string {
	if (len(conf.MongoDB["login"]) == 0) {
		conf.MongoDB["login"] = "Guest"
	}
	if (len(conf.MongoDB["password"]) == 0) {
		conf.MongoDB["password"] = "Guest"
	}
	if (len(conf.MongoDB["hostname"]) == 0) {
		conf.MongoDB["hostname"] = "localhost"
	}
	if (len(conf.MongoDB["port"]) == 0) {
		conf.MongoDB["port"] = "27017"
	}
	if (len(conf.MongoDB["database"]) == 0) {
		conf.MongoDB["database"] = "PaintLayering"
	}
	if (len(conf.MongoDB["options"]) == 0) {
		conf.MongoDB["options"] = ""
	}
	conn := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?%s", conf.MongoDB["login"], conf.MongoDB["password"], conf.MongoDB["hostname"], conf.MongoDB["port"], conf.MongoDB["database"], conf.MongoDB["options"])
	
	return conn
}

func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
	handleError(err)
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	var filename string
	var conf Config
	var e []envId
	
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
	
	if (len(conf.MongoDB["description"]) != 0) {
		log.Println(conf.MongoDB["description"])		
	}
	
	if (len(conf.MongoDB["timeout"]) == 0) {
		conf.MongoDB["timeout"] = "15"
	}
	
	timeint, err := strconv.Atoi(conf.MongoDB["timeout"])
	handleError(err)

	timeout := time.Duration(timeint)
	
	session, err := mgo.DialWithTimeout(url, timeout * time.Second)
	handleError(err)

	defer session.Close()
	
	session.SetMode(mgo.Monotonic, true)
	
	err = getJson(conf.Url["Layer"], &e)
	handleError(err)
	
	fmt.Printf("Id: %v\n", e[0].Skus[0].Id)
	fmt.Printf("ProductName: %v\n", e[0].Skus[0].Title)
	
	os.Exit(0)
}	