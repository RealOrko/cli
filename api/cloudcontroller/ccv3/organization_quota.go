package ccv3

import (
	"github.com/realorko/cli/api/cloudcontroller/ccv3/public"
)

// OrganizationQuota represents a Cloud Controller organization quota.
type OrganizationQuota struct {
	Quota
}

func (client *Client) ApplyOrganizationQuota(quotaGuid, orgGuid string) (RelationshipList, Warnings, error) {
	var responseBody RelationshipList

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PostOrganizationQuotaApplyRequest,
		URIParams:    internal.Params{"quota_guid": quotaGuid},
		RequestBody:  RelationshipList{GUIDs: []string{orgGuid}},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

func (client *Client) CreateOrganizationQuota(orgQuota OrganizationQuota) (OrganizationQuota, Warnings, error) {
	var responseOrgQuota OrganizationQuota

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PostOrganizationQuotaRequest,
		RequestBody:  orgQuota,
		ResponseBody: &responseOrgQuota,
	})

	return responseOrgQuota, warnings, err
}

func (client *Client) DeleteOrganizationQuota(quotaGUID string) (JobURL, Warnings, error) {
	jobURL, warnings, err := client.MakeRequest(RequestParams{
		RequestName: internal.DeleteOrganizationQuotaRequest,
		URIParams:   internal.Params{"quota_guid": quotaGUID},
	})

	return jobURL, warnings, err
}

func (client *Client) GetOrganizationQuota(quotaGUID string) (OrganizationQuota, Warnings, error) {
	var responseBody OrganizationQuota

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.GetOrganizationQuotaRequest,
		URIParams:    internal.Params{"quota_guid": quotaGUID},
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}

func (client *Client) GetOrganizationQuotas(query ...Query) ([]OrganizationQuota, Warnings, error) {
	var resources []OrganizationQuota

	_, warnings, err := client.MakeListRequest(RequestParams{
		RequestName:  internal.GetOrganizationQuotasRequest,
		Query:        query,
		ResponseBody: OrganizationQuota{},
		AppendToList: func(item interface{}) error {
			resources = append(resources, item.(OrganizationQuota))
			return nil
		},
	})

	return resources, warnings, err
}

func (client *Client) UpdateOrganizationQuota(orgQuota OrganizationQuota) (OrganizationQuota, Warnings, error) {
	orgQuotaGUID := orgQuota.GUID
	orgQuota.GUID = ""

	var responseBody OrganizationQuota

	_, warnings, err := client.MakeRequest(RequestParams{
		RequestName:  internal.PatchOrganizationQuotaRequest,
		URIParams:    internal.Params{"quota_guid": orgQuotaGUID},
		RequestBody:  orgQuota,
		ResponseBody: &responseBody,
	})

	return responseBody, warnings, err
}
