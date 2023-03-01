package com

import (
	"encoding/json"

	"example.local/hpecom/rest"
)

type JobList struct {
	Offset int   `json:"offset,omitempty"`
	Count  int   `jsont:"count,omitempty"`
	Total  int   `json:"total,omitempty"`
	Items  []Job `json:"items,omitempty"`
}

type Job struct {
	Id                    string      `json:"id,omitempty"`
	ParentJobId           string      `json:"parentJobId,omitempty"`
	Type                  string      `json:"type,omitempty"`
	ResourceUri           string      `json:"resourceUri,omitempty"`
	Name                  string      `json:"name,omitempty"`
	Generation            int         `json:"generation,omitempty"`
	CreatedAt             string      `json:"createdAt,omitempty"`
	UpdatedAt             string      `json:"updatedAt,omitempty"`
	JobTemplateUri        string      `json:"jobTemplateUri,omitempty"`
	AssociatedResourceUri string      `json:"associatedResourceUri,omitempty"`
	Resource              JobResource `json:"resource,omitempty"`
	Data                  interface{} `json:"data,omitempty"`
	State                 string      `json:"state,omitempty"`
	ResultCode            string      `json:"resultCode,omitempty"`
	Status                string      `json:"status,omitempty"`
	StatusDetails         struct{}    `json:"statusDetails,omitempty"`
}

type JobResource struct {
	ResourceUri string `json:"resourceUri,omitempty"`
	Type        string `json:"type,omitempty"`
}

type JobData struct {
}

type DataResource struct {
	Id                     string                `json:"id,omitempty"`
	Name                   string                `json:"name,omitempty"`
	Customer_id            string                `json:"customer_id,omitempty"`
	Devices                []DataResourceDevices `json:"devices,omitempty"`
	Auto_fw_update_enabled string                `json:"auto_fw_update_enabled,omitempty"` //should probably be bool, but API returns null string
	Auto_fw_update_on_add  bool                  `json:"auto_fw_on_add,omitempty"`
	Install_sw_drivers     string                `json:"install_sw_drivers,omitempty"` //should probably be bool, but API returns null string
	Server_settings_uris   []string              `json:"server_settings_uris,omitempty"`
	Server_policies        ServerPolicies        `json:"server_policies,omitempty"`
}

type DataResourceDevices struct {
	Serial             string `json:"serial,omitempty"`
	Product_id         string `json:"product_id,omitempty"`
	Group_id           string `json:"group_id,omitempty"`
	State              string `json:"state,omitempty"`
	Server_id          string `json:"server_id,omitempty"`
	Server_uri         string `json:"server_uri,omitempty"`
	Subscription_state string `json:"subscription_state,omitempty"`
	Subscription_tier  string `json:"subscription_tier,omitempty"`
}

type ServerPolicies struct {
	On_server_add string `json:"on_server_add,omitempty"`
}

var (
	jobsUri = "/compute-ops/v1beta3/jobs"
)

func (c *ComClient) GetJobs(filters []string, offset, limit string) (JobList, error) {

	var (
		q    map[string]interface{}
		jobs JobList
	)
	c.SetAuthHeaderOptions(c.GetAuthHeaders())

	q = c.SetQueryParams(filters, offset, limit)

	data, err := c.RestAPICall(rest.GET, jobsUri, nil, q)
	if err != nil {
		return jobs, err
	}

	if err := json.Unmarshal([]byte(data), &jobs); err != nil {
		return jobs, err
	}

	return jobs, nil
}
