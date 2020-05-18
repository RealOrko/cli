package ccv3

import (
	"encoding/json"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"github.com/realorko/cli/api/cloudcontroller/ccv3/public"
)

// Relationship represents a one to one relationship.
// An empty GUID will be marshaled as `null`.
type Relationship struct {
	GUID string
}

func (r Relationship) MarshalJSON() ([]byte, error) {
	if r.GUID == "" {
		var emptyCCRelationship struct {
			Data interface{} `json:"data"`
		}
		return json.Marshal(emptyCCRelationship)
	}

	var ccRelationship struct {
		Data struct {
			GUID string `json:"guid"`
		} `json:"data"`
	}

	ccRelationship.Data.GUID = r.GUID
	return json.Marshal(ccRelationship)
}

func (r *Relationship) UnmarshalJSON(data []byte) error {
	var ccRelationship struct {
		Data struct {
			GUID string `json:"guid"`
		} `json:"data"`
	}

	err := cloudcontroller.DecodeJSON(data, &ccRelationship)
	if err != nil {
		return err
	}

	r.GUID = ccRelationship.Data.GUID
	return nil
}

// DeleteIsolationSegmentOrganization will delete the relationship between
// the isolation segment and the organization provided.
func (client *Client) DeleteIsolationSegmentOrganization(isolationSegmentGUID string, orgGUID string) (Warnings, error) {
	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.DeleteIsolationSegmentRelationshipOrganizationRequest,
		URIParams:   internal.Params{"isolation_segment_guid": isolationSegmentGUID, "organization_guid": orgGUID},
	})

	return warnings, err
}

// DeleteServiceInstanceRelationshipsSharedSpace will delete the sharing relationship
// between the service instance and the shared-to space provided.
func (client *Client) DeleteServiceInstanceRelationshipsSharedSpace(serviceInstanceGUID string, spaceGUID string) (Warnings, error) {
	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.DeleteServiceInstanceRelationshipsSharedSpaceRequest,
		URIParams:   internal.Params{"service_instance_guid": serviceInstanceGUID, "space_guid": spaceGUID},
	})

	return warnings, err
}

// GetOrganizationDefaultIsolationSegment returns the relationship between an
// organization and it's default isolation segment.
func (client *Client) GetOrganizationDefaultIsolationSegment(orgGUID string) (Relationship, Warnings, error) {
	var responseBody Relationship

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetOrganizationRelationshipDefaultIsolationSegmentRequest,
		URIParams:    internal.Params{"organization_guid": orgGUID},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

// GetSpaceIsolationSegment returns the relationship between a space and it's
// isolation segment.
func (client *Client) GetSpaceIsolationSegment(spaceGUID string) (Relationship, Warnings, error) {
	var responseBody Relationship

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetSpaceRelationshipIsolationSegmentRequest,
		URIParams:    internal.Params{"space_guid": spaceGUID},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

// SetApplicationDroplet sets the specified droplet on the given application.
func (client *Client) SetApplicationDroplet(appGUID string, dropletGUID string) (Relationship, Warnings, error) {
	var responseBody Relationship

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PatchApplicationCurrentDropletRequest,
		URIParams:    internal.Params{"app_guid": appGUID},
		RequestBody:  Relationship{GUID: dropletGUID},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

// UpdateOrganizationDefaultIsolationSegmentRelationship sets the default isolation segment
// for an organization on the controller.
// If isoSegGuid is empty it will reset the default isolation segment.
func (client *Client) UpdateOrganizationDefaultIsolationSegmentRelationship(orgGUID string, isoSegGUID string) (Relationship, Warnings, error) {
	var responseBody Relationship

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PatchOrganizationRelationshipDefaultIsolationSegmentRequest,
		URIParams:    internal.Params{"organization_guid": orgGUID},
		RequestBody:  Relationship{GUID: isoSegGUID},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

// UpdateSpaceIsolationSegmentRelationship assigns an isolation segment to a space and
// returns the relationship.
func (client *Client) UpdateSpaceIsolationSegmentRelationship(spaceGUID string, isolationSegmentGUID string) (Relationship, Warnings, error) {
	var responseBody Relationship

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PatchSpaceRelationshipIsolationSegmentRequest,
		URIParams:    internal.Params{"space_guid": spaceGUID},
		RequestBody:  Relationship{GUID: isolationSegmentGUID},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}
