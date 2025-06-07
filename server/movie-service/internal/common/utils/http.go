package utils

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetProxy() (error, *http.Transport) {
	proxy, err := url.Parse("http://" + os.Getenv("PROXY_USERNAME") + ":" + os.Getenv("PROXY_PASSWORD") + "@" + os.Getenv("PROXY_ADDRESS") + ":" + os.Getenv("PROXY_PORT"))
	if err != nil {
		return fmt.Errorf("invalid proxy URL"), nil
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
		DialContext: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).DialContext,
	}

	return nil, transport
}
