package cmd

import (
	"net/http"
	"os"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

// GetAuthenticatedClient test
func GetAuthenticatedClient(transport *http.Transport) (*gophercloud.ServiceClient, error) {
	// Get authentication details from env vars
	auth, err := openstack.AuthOptionsFromEnv()
	auth.TenantID = os.Getenv("OS_PROJECT_ID")

	if err != nil {
		return nil, err
	}

	client, err := openstack.NewClient(auth.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	// config := &tls.Config{InsecureSkipVerify: true}
	// transport := &http.Transport{TLSClientConfig: config}
	client.HTTPClient.Transport = transport

	err = openstack.AuthenticateV3(client, auth)
	if err != nil {
		return nil, err
	}

	serviceClient, err := openstack.NewComputeV2(client, gophercloud.EndpointOpts{
		Region: "au-east-2",
	})
	if err != nil {
		return nil, err
	}

	return serviceClient, nil
}
