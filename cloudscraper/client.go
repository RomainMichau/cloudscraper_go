package cloudscraper

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

type CloudScrapper struct {
	client        cycletls.CycleTLS
	defaultHeader map[string]string
	ja3           string
	userAgent     string
}

func (cs CloudScrapper) Do(url string, options cycletls.Options, method string) (cycletls.Response, error) {
	for k, v := range cs.defaultHeader {
		options.Headers[k] = v
	}
	if options.UserAgent == "" {
		options.UserAgent = cs.userAgent
	}
	if options.Ja3 == "" {
		options.Ja3 = cs.ja3
	}
	return cs.client.Do(url, options, method)
}

func (cs CloudScrapper) Get(url string, headers map[string]string, body string) (cycletls.Response, error) {
	options := cycletls.Options{
		Ja3:       cs.ja3,
		Body:      body,
		UserAgent: cs.userAgent,
		Headers:   headers}
	return cs.Do(url, options, "GET")
}

func (cs CloudScrapper) Post(url string, headers map[string]string, body string) (cycletls.Response, error) {
	options := cycletls.Options{
		Ja3:       cs.ja3,
		Body:      body,
		UserAgent: cs.userAgent,
		Headers:   headers}
	return cs.Do(url, options, "POST")
}

func Init(mobile bool) *CloudScrapper {
	browserConf, _ := GetUserAgents(mobile)
	cycleTstClient := cycletls.Init()
	p := CloudScrapper{client: cycleTstClient, defaultHeader: browserConf.Headers, ja3: browserConf.Ja3, userAgent: browserConf.UserAgent}
	return &p
}