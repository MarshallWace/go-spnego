package spnego

import (
	"net/http"
	"net/url"
	"testing"
)

func TestGetHostnameEmptyHost(t *testing.T) {
	desired := "example.com"

	url, err := url.Parse("http://example.com/path")
	if err != nil {
		panic(err)
	}
	req := &http.Request{
		URL: url,
	}
	t.Logf("url %s, host %s\n", req.URL.Host, req.Host)

	h, err := getHostname(req, false)
	if err != nil {
		panic(err)
	}

	if h != desired {
		t.Errorf("Hostname was incorrect: wanted %s, got %s.", desired, h)
	}
}

func TestGetHostnameEmptyHostPort(t *testing.T) {
	desired := "example.com:8080"

	url, err := url.Parse("http://example.com:8080/path")
	if err != nil {
		panic(err)
	}
	req := &http.Request{
		URL: url,
	}
	t.Logf("url %s, host %s\n", req.URL.Host, req.Host)

	h, err := getHostname(req, false)
	if err != nil {
		panic(err)
	}

	if h != desired {
		t.Errorf("Hostname was incorrect: wanted %s, got %s.", desired, h)
	}
}

func TestGetHostnameHost(t *testing.T) {
	desired := "example.com"

	req, err := http.NewRequest("GET", "http://example.com/path", nil)
	if err != nil {
		panic(err)
	}
	t.Logf("url %s, host %s\n", req.URL.Host, req.Host)

	h, err := getHostname(req, false)
	if err != nil {
		panic(err)
	}

	if h != desired {
		t.Errorf("Hostname was incorrect: wanted %s, got %s.", desired, h)
	}
}

func TestGetHostnameHostWithPort(t *testing.T) {
	desired := "example.com:8080"

	req, err := http.NewRequest("GET", "http://example.com:8080/path", nil)
	if err != nil {
		panic(err)
	}
	t.Logf("url %s, host %s\n", req.URL.Host, req.Host)

	h, err := getHostname(req, false)
	if err != nil {
		panic(err)
	}

	if h != desired {
		t.Errorf("Hostname was incorrect: wanted %s, got %s.", desired, h)
	}
}

func TestGetHostnameOverride(t *testing.T) {
	desired := "override.org"

	req, err := http.NewRequest("GET", "http://example.com:8080/path", nil)
	if err != nil {
		panic(err)
	}
	req.Host = "override.org"
	t.Logf("url %s, host %s\n", req.URL.Host, req.Host)

	h, err := getHostname(req, false)
	if err != nil {
		panic(err)
	}

	if h != desired {
		t.Errorf("Hostname was incorrect: wanted %s, got %s.", desired, h)
	}
}

func TestGetHostnameOverrideWithPort(t *testing.T) {
	desired := "override.org:1234"

	req, err := http.NewRequest("GET", "http://example.com:8080/path", nil)
	if err != nil {
		panic(err)
	}
	req.Host = "override.org:1234"
	t.Logf("url %s, host %s\n", req.URL.Host, req.Host)

	h, err := getHostname(req, false)
	if err != nil {
		panic(err)
	}

	if h != desired {
		t.Errorf("Hostname was incorrect: wanted %s, got %s.", desired, h)
	}
}
