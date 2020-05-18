package ccv3

import (
	"github.com/realorko/cli/api/cloudcontroller/ccv3/public"
)

// GetApplicationManifest returns a (YAML) manifest for an application and its
// underlying processes.
func (client *Client) GetApplicationManifest(appGUID string) ([]byte, Warnings, error) {
	bytes, warnings, err := client.MakeRequestReceiveRaw(
		internal.GetApplicationManifestRequest,
		internal.Params{"app_guid": appGUID},
		"application/x-yaml",
	)

	return bytes, warnings, err
}
