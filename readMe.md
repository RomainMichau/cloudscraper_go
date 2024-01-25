# CloudScraper Go

Simple go module to bypass CloudFlare securities

For now only passive checks based on TLS footprint (JA3) and user agents are handled  
Inspired of python project https://github.com/venomous/cloudscraper

Rely on [CycleTls](https://github.com/Danny-Dasilva/CycleTLS) to customize JA3

JA3 list:
https://raw.githubusercontent.com/trisulnsm/trisul-scripts/master/lua/frontend_scripts/reassembly/ja3/prints/ja3fingerprint.json

Get my JA3: https://kawayiyi.com/tls

# Example of usage
## Simple Post
```go
package main

import (
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

func main() {

	client, _ := cloudscraper.Init(false, false)
	res, _ := client.Post("https://www.facebook.com/anything/", make(map[string]string), "")
	res2, _ := client.Get("https://www.facebook.com/anything/", make(map[string]string), "")
	print(res.Body)
	print(res2.Body)
}
```
## Customize request
```go
package main

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

func main() {

	client, _ := cloudscraper.Init(false, false)
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
## Async Request
```go
package main

import (
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/RomainMichau/cloudscraper_go/cloudscraper"
)

func main() {

	client, _ := cloudscraper.Init(false, true)
	options := cycletls.Options{
		Headers:         map[string]string{"my_custom_header": "header_value"},
		Body:            "",
		Timeout:         10,
		DisableRedirect: true,
	}
	client.Queue("https://www.facebook.com/anything", options, "PUT")
	respChannel := client.RespChan()
	res := <-respChannel
	print(res.Body)
}

```

# Todo
- Bypass cloudflare active counter-measure (JS challenge) as done on python [cloudscraper](https://github.com/venomous/cloudscraper) 
