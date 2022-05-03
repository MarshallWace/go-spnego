// SPDX-License-Identifier: MIT

package spnego

import (
	"net"
	"net/http"
	"strings"
)

// Provider is the interface that wraps OS agnostic functions for handling SPNEGO communication
type Provider interface {
	SetSPNEGOHeader(*http.Request, bool) error
}

func canonicalizeHostname(hostname string) (string, error) {
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return "", err
	}
	if len(addrs) < 1 {
		return hostname, nil
	}

	names, err := net.LookupAddr(addrs[0])
	if err != nil {
		return "", err
	}
	if len(names) < 1 {
		return hostname, nil
	}

	return strings.TrimRight(names[0], "."), nil
}

func getHostname(req *http.Request, canonicalize bool) (string, error) {
	var err error

	h := req.URL.Host // SPN should contain the port, if non-standard (https://social.technet.microsoft.com/wiki/contents/articles/717.service-principal-names-spn-setspn-syntax.aspx)
	if req.Host != "" {
		h = req.Host
	}
	if canonicalize {
		if h, err = canonicalizeHostname(h); err != nil {
			return "", err
		}
	}

	return h, nil
}
