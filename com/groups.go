package com

import (
	"encoding/json"

	"example.local/hpecom/rest"
)

type GroupList struct {
	Offset int      `json:"offset,omitempty"`
	Count  int      `jsont:"count,omitempty"`
	Total  int      `json:"total,omitempty"`
	Items  []Groups `json:"items,omitempty"`
}

type Groups struct {
	Type                  string                 `json:"type,omitempty"`
	ResourceURI           string                 `json:"resourceUri,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	Generation            int                    `json:"generation,omitempty"`
	CreatedAt             string                 `json:"createdAt,omitempty"`
	UpdatedAt             string                 `json:"updatedAt,omitempty"`
	Description           string                 `json:"description,omitempty"`
	AutoFwUpdateOnAdd     bool                   `json:"autoFwUpdateOnAdd,omitempty"`
	GroupComplianceStatus string                 `json:"groupComplianceStatus,omitempty"`
	ServerSettingsUris    []string               `json:"serverSettingsUris,omitempty"`
	PlatformFamily        string                 `json:"platformFamily,omitempty"`
	DevicesURI            string                 `json:"devicesUri,omitempty"`
	Devices               []GroupDevices         `json:"devices,omitempty"`
	ServerPolicies        GroupPolicies          `json:"serverPolicies,omitempty"`
	AutoAddServerTags     map[string]interface{} `json:"autoAddServerTags,omitempty"`
	GroupMeta             map[string]interface{} `json:"groupMeta,omitempty"`
	Client                *ComClient             `json:"-"`
}

type GroupDevices struct {
	Serial            string `json:"serial,omitempty"`
	ProductId         string `json:"productId,omitempty"`
	ETag              string `json:"etag,omitempty"`
	ServerId          string `json:"serverId,omitempty"`
	ID                string `json:"id,omitempty"`
	ResourceUri       string `json:"resourceUri,omitempty"`
	Type              string `json:"type,omitempty"`
	ServerUri         string `json:"serverUri,omitempty"`
	State             string `json:"state,omitempty"`
	GroupId           string `json:"groupId,omitempty"`
	SubscriptionState string `json:"subscriptionState,omitempty"`
	SubscriptionTier  string `json:"subscriptionTier,omitempty"`
}

type ServerAddPolicies struct {
	FirmwareDowngrade     bool   `json:"firmwareDowngrade,omitempty"`
	FirmwareUpdate        bool   `json:"firmwareUpdate,omitempty"`
	BiosApplySettings     bool   `json:"biosApplySettings,omitempty"`
	StorageConfiguration  bool   `json:"storageConfiguration,omitempty"`
	StorageVolumeDeletion bool   `json:"storageVolumeDeletion,omitempty"`
	StorageVolumeName     string `json:"storageVolumeName,omitempty"`
}

type GroupPolicies struct {
	OnServerAdd ServerAddPolicies `json:"onServerAdd,omitempty"`
}

var (
	groupsUri = "/compute-ops/v1beta2/groups"
)

func (c *ComClient) GetGroups(filters []string, offset, limit string) (GroupList, error) {

	var (
		groups GroupList
		q      map[string]interface{}
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaders())

	q = c.SetQueryParams(filters, offset, limit)

	data, err := c.RestAPICall(rest.GET, groupsUri, nil, q)
	if err != nil {
		return groups, err
	}

	if err := json.Unmarshal([]byte(data), &groups); err != nil {
		return groups, err
	}

	for i := 0; i < groups.Count; i++ {
		groups.Items[i].Client = c
	}
	return groups, nil
}
