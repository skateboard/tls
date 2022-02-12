package client

import utls "github.com/refraction-networking/utls"

type Browser struct {
	UserAgent 		string
	ClientID utls.ClientHelloID
}

var Browsers = map[string]Browser{
	"chrome_58": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		ClientID: utls.HelloChrome_58,
	},
	"chrome_62": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		ClientID: utls.HelloChrome_62,
	},
	"chrome_70": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36",
		ClientID: utls.HelloChrome_70,
	},
	"chrome_72": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3282.140 Safari/537.36",
		ClientID: utls.HelloChrome_72,
	},
	"chrome_83": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
		ClientID: utls.HelloChrome_83,
	},
	"firefox_55": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:55.0) Gecko/20100101 Firefox/55.0",
		ClientID: utls.HelloFirefox_55,
	},
	"firefox_56": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:60.0) Gecko/20100101 Firefox/56.0",
		ClientID: utls.HelloFirefox_56,
	},
	"firefox_63": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:61.0) Gecko/20100101 Firefox/63.0",
		ClientID: utls.HelloFirefox_63,
	},
	"firefox_65": {
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:64.0) Gecko/20100101 Firefox/65.0",
		ClientID: utls.HelloFirefox_65,
	},
	"ios_11": {
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		ClientID: utls.HelloIOS_11_1,
	},
	"ios_12": {
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1",
		ClientID: utls.HelloIOS_12_1,
	},
}