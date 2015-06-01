package auth

import (
	"fmt"
	"os"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
)

// NewClient creates and returns a Rackspace client for the given service.
func NewClient(t string) *gophercloud.ServiceClient {
	var err error
	ao, err := authMethod()
	if err != nil {
		fmt.Printf("Error building AuthOptions: %s\n", err)
		os.Exit(1)
	}
	pc, err := rackspace.AuthenticatedClient(ao)
	if err != nil {
		fmt.Printf("Error creating ProviderClient: %s\n", err)
		os.Exit(1)
	}
	var sc *gophercloud.ServiceClient
	switch t {
	case "compute":
		sc, err = rackspace.NewComputeV2(pc, gophercloud.EndpointOpts{
			Region: os.Getenv("RS_REGION_NAME"),
		})
		break
	case "blockstorage":
		sc, err = rackspace.NewBlockStorageV1(pc, gophercloud.EndpointOpts{
			Region: os.Getenv("RS_REGION_NAME"),
		})
		break
	case "networking":
		sc, err = rackspace.NewNetworkV2(pc, gophercloud.EndpointOpts{
			Region: os.Getenv("RS_REGION_NAME"),
		})
		break
	}
	if err != nil {
		fmt.Printf("Error creating ServiceClient (%s): %s\n", err, t)
		os.Exit(1)
	}
	// sc.UserAgent.Prepend("rackcli/" + util.Version)
	return sc
}

// authMethod determines the appropriate authentication method for the user.
func authMethod() (gophercloud.AuthOptions, error) {
	return envvars()
}