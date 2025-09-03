package cloudscraper

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strings"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
)

const JA3_FINGERPRINT_URL = "https://raw.githubusercontent.com/trisulnsm/trisul-scripts/master/lua/frontend_scripts/reassembly/ja3/prints/ja3fingerprint.json"

type browserDescription struct {
	Headers map[string]map[string]string `json:"headers"`
}

type BrowserConf struct {
	UserAgent string
	Ja3       string
	Headers   map[string]string
}

//go:embed resources/browsers.json
var browsersJson string

func readJsonFile() (browserDescription, error) {
	// Open our jsonFile
	var browsers browserDescription
	err := json.Unmarshal([]byte(browsersJson), &browsers)
	// defer the closing of our jsonFile so that we can parse it later on
	return browsers, err
}

type Device string

const (
	Chrome  Device = "chrome"
	Firefox Device = "firefox"
	Safari  Device = "safari"
	Mobile  Device = "mobile"
	Default Device = "default"
)

func getUserAgents(device Device) (BrowserConf, error) {
	var randUA string

	switch device {
	case Chrome:
		randUA = browser.Chrome()
	case Firefox:
		randUA = browser.Firefox()
	case Safari:
		randUA = browser.Safari()
	case Mobile:
		randUA = browser.Mobile()
	default:
		randUA = browser.Random()
	}

	browsersDescription, err := readJsonFile()
	if err != nil {
		return BrowserConf{}, err
	}

	ja3, err := pickJA3(device)
	if err != nil {
		return BrowserConf{}, err
	}

	return BrowserConf{UserAgent: randUA, Ja3: ja3, Headers: browsersDescription.Headers[string(device)]}, nil
}

// fetched from url
var recourseJA3 string

func fetchJA3() error {
	cl := &http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	resp, err := cl.Get(JA3_FINGERPRINT_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	recourseJA3 = buf.String()
	return nil
}

var ja3Map = map[Device][]string{
	Chrome:  {"Chrome", "Chromium"},
	Firefox: {"Firefox"},
	Safari:  {"Safari"},
	Mobile:  {"Android", "iOS"},
	Default: {"Chrome", "Safari", "Firefox"},
}

func pickJA3(device Device) (string, error) {
	lines := strings.Split(recourseJA3, "\n")

	var ja3Candidates []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "{") {
			continue
		}

		var ja3 struct {
			Desc string `json:"desc"`
			Str  string `json:"ja3_str"`
		}
		if err := json.Unmarshal([]byte(line), &ja3); err != nil {
			continue
		}

		for _, kw := range ja3Map[device] {
			if strings.Contains(strings.ToLower(ja3.Desc), strings.ToLower(kw)) {
				ja3Candidates = append(ja3Candidates, ja3.Str)
			}
		}
	}

	if len(ja3Candidates) == 0 {
		return "", errors.New("no ja3 match for device")
	}

	return ja3Candidates[rInt(len(ja3Candidates))], nil
}

// better random
func rInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}
