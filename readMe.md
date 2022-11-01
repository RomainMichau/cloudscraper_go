# CloudScraper Go

Simple go module to bypass CloudFlare securities

For now only passive checks based on TLS footprint (JA3) and user agents are handled  
Inspired of python project https://github.com/venomous/cloudscraper

Rely on [CycleTls](https://github.com/Danny-Dasilva/CycleTLS) to customize JA3

JA3 list:
https://raw.githubusercontent.com/trisulnsm/trisul-scripts/master/lua/frontend_scripts/reassembly/ja3/prints/ja3fingerprint.json

# Example of usage

```go
package main

import (
	"github.com/RomainMichau/cloud_scraper_go/cloudscrapper"
)

func main() {

	client := cloudscrapper.Init(false)
	res, _ := client.Post("https://www.facebook.com/anything/", make(map[string]string), "")
	print(res.Body)
}
```

```go
package main

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/RomainMichau/cloud_scraper_go/cloudscrapper"
)

func main() {

	client := cloudscrapper.Init(false)
	options := cycletls.Options{
		Headers:         map[string]string{"my_custom_header": "header_value"},
		Body:            "",
		Proxy:           "proxy.company.com",
		Timeout:         10,
		DisableRedirect: true,
	}
	res, _ := client.Do("https://www.facebook.com/anything", options, "PUT")
	print(res.Body)
}
```

# Todo
- Bypass cloudflare active counter-measure (JS challenge) as done on python [cloudscraper](https://github.com/venomous/cloudscraper) 