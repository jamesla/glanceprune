package cmd

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSum(t *testing.T) {

	os.Setenv("OS_AUTH_URL", "https://127.0.0.1")
	os.Setenv("OS_USERNAME", "user")
	os.Setenv("OS_PASSWORD", "pass")
	os.Setenv("OS_DOMAIN_NAME", "domain.local")

	// Stub endpoint
	// config := &tls.Config{InsecureSkipVerify: true}
	// transport := &http.Transport{TLSClientConfig: config}
	s := httptest.NewTLSServer(handler)

	transport := &http.Transport{
		DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
			return net.Dial(network, s.Listener.Addr().String())
		},
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Get an Authenticated client
	_, err := GetAuthenticatedClient(transport)
	if err != nil {
		t.Errorf("Auth client failed: %s", err)
		return
	}
}
