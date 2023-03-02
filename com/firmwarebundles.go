package com

import (
	"encoding/json"

	"github.com/ejschulte/hpecom-golang/rest"
)

type FirmwareBundleList struct {
	Offset int              `json:"offset,omitempty"`
	Count  int              `jsont:"count,omitempty"`
	Total  int              `json:"total,omitempty"`
	Items  []FirmwareBundle `json:"items,omitempty"`
}

type FirmwareBundle struct {
	Id              string      `json:"id,omitempty"`
	Type            string      `json:"type,omitempty"`
	ResourceUri     string      `json:"resourceUri,omitempty"`
	Generation      int         `json:"generation,omitempty"`
	CreatedAt       string      `json:"createdAt,omitempty"`
	UpdatedAt       string      `json:"updatedAt,omitempty"`
	Name            string      `json:"name,omitempty"`
	Description     string      `json:"description,omitempty"`
	ReleaseDate     string      `json:"releaseDate,omitempty"`
	ReleaseVersion  string      `json:"releaseVersion,omitempty"`
	ReleaseNotes    string      `json:"releaseNotes,omitempty"`
	SupportUrl      string      `json:"supportUrl,omitempty"`
	Enhancements    string      `json:"enhancements,omitempty"`
	Advisories      string      `json:"advisories,omitempty"`
	SupportedOsList []string    `json:"supportedOsList,omitempty"`
	IsActive        bool        `json:"isActive,omitempty"`
	Summary         string      `json:"summary,omitempty"`
	BundleType      string      `json:"bundleType,omitempty"`
	HotfixBaseUri   string      `json:"hotfixBaseUri,omitempty"`
	VmwareAddonInfo interface{} `json:"vmwareAddonInfo,omitempty"` // not currently defined in API
}

var (
	bundleUri = "/compute-ops/v1beta2/firmware-bundles"
)

func (c *ComClient) GetFirmwareBundles(filters []string, offset, limit string) (FirmwareBundleList, error) {
	var (
		q       map[string]interface{}
		bundles FirmwareBundleList
	)
	c.SetAuthHeaderOptions(c.GetAuthHeaders())
	q = c.SetQueryParams(filters, offset, limit)
	data, err := c.RestAPICall(rest.GET, bundleUri, nil, q)
	if err != nil {
		return bundles, err
	}
	if err := json.Unmarshal([]byte(data), &bundles); err != nil {
		return bundles, err
	}
	return bundles, nil
}
