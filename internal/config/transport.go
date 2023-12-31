// Copyright © 2018 Heptio
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"net/http"
	"time"
)

// TransportConfig describes a configured httpClient
type TransportConfig struct {
	HTTPClient *http.Client
}

// NewTransportConfig returns a TransportConfig with configured httpClient
func NewTransportConfig(trustedCA []byte) *TransportConfig {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	if len(trustedCA) > 0 {
		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(trustedCA); !ok {
			log.Println("No certs appended, using system certs only")
		}
	}

	// Transport based on http.DefaultTransport
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			RootCAs:    rootCAs,
			MinVersion: tls.VersionTLS12,
		},
	}

	httpClient := &http.Client{
		Transport: t,
	}

	return &TransportConfig{
		HTTPClient: httpClient,
	}
}
