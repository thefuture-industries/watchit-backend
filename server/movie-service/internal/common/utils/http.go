package utils

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetProxy() (*http.Transport, error) {
	proxy, err := url.Parse("http://" + os.Getenv("PROXY_USERNAME") + ":" + os.Getenv("PROXY_PASSWORD") + "@" + os.Getenv("PROXY_ADDRESS") + ":" + os.Getenv("PROXY_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid proxy URL")
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
		DialContext: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).DialContext,
	}

	return transport, nil
}
