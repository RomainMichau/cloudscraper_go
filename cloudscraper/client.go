package cloudscraper

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

type CloudScrapper struct {
	client        cycletls.CycleTLS
	respChan      chan cycletls.Response
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

// Queue Push a request in the Queue
// Read the result with RespChan()
func (cs CloudScrapper) Queue(url string, options cycletls.Options, method string) {
	for k, v := range cs.defaultHeader {
		options.Headers[k] = v
	}
	if options.UserAgent == "" {
		options.UserAgent = cs.userAgent
	}
	if options.Ja3 == "" {
		options.Ja3 = cs.ja3
	}

	cs.client.Queue(url, options, method)
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

func (cs CloudScrapper) RespChan() chan cycletls.Response {
	return cs.respChan
}

func Init(mobile bool, workers bool) (*CloudScrapper, error) {
	browserConf, err := getUserAgents(mobile)
	if err != nil {
		return nil, err
	}
	cycleTstClient := cycletls.Init(workers)
	p := CloudScrapper{client: cycleTstClient, defaultHeader: browserConf.Headers, ja3: browserConf.Ja3, userAgent: browserConf.UserAgent,
		respChan: cycleTstClient.RespChan}
	return &p, nil
}
