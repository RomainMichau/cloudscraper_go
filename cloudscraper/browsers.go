package cloudscraper

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type BrowserDescription struct {
	UserAgents UserAgents                   `json:"user_agents"`
	Ja3        map[string]string            `json:"ja3"`
	Headers    map[string]map[string]string `json:"headers"`
}

type UserAgents struct {
	// os -> browser - [user-agents]
	Desktop map[string]map[string][]string `json:"desktop"`
	Mobile  map[string]map[string][]string `json:"mobile"`
}

type BrowserConf struct {
	UserAgent string
	Ja3       string
	Headers   map[string]string
}

func readJsonFile() (BrowserDescription, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("ressources/browsers.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	byteValue, _ := io.ReadAll(jsonFile)
	var browsers BrowserDescription
	err = json.Unmarshal(byteValue, &browsers)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return browsers, err
}

func GetUserAgents(mobile bool) (BrowserConf, error) {
	rand.Seed(time.Now().UnixNano())
	var userAgents map[string]map[string][]string
	browsersDescription, err := readJsonFile()
	if err != nil {
		return BrowserConf{}, err
	}
	if mobile {
		userAgents = browsersDescription.UserAgents.Mobile
	} else {
		userAgents = browsersDescription.UserAgents.Desktop
	}
	var osList []string
	for k := range userAgents {
		osList = append(osList, k)
	}
	rnd := rand.Intn(len(osList))
	pickedOs := userAgents[osList[rnd]]
	var browserList []string
	for k := range pickedOs {
		browserList = append(browserList, k)
	}
	rnd = rand.Intn(len(browserList))
	browserName := browserList[rnd]
	pickedBrowser := pickedOs[browserName]
	rnd = rand.Intn(len(pickedBrowser))
	pickedUserAgent := pickedBrowser[rnd]
	ja3 := browsersDescription.Ja3[browserName]
	return BrowserConf{UserAgent: pickedUserAgent, Ja3: ja3, Headers: browsersDescription.Headers[browserName]}, nil
}
