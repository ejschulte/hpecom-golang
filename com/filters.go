package com

import (
	"encoding/json"

	"github.com/ejschulte/hpecom-golang/rest"
)

type FilterList struct {
	Offset int      `json:"offset,omitempty"`
	Count  int      `jsont:"count,omitempty"`
	Total  int      `json:"total,omitempty"`
	Items  []Filter `json:"items,omitempty"`
}

type Filter struct {
	Id                 string `json:"id,omitempty"`
	Type               string `json:"type,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	CreatedAt          string `json:"createdAt,omitempty"`
	UpdatedAt          string `json:"updatedAt,omitempty"`
	ResourceUri        string `json:"resourceUri,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	ReadOnly           bool   `json:"readOnly,omitempty"`
	FilterResourceType string `json:"filterResourceType,omitempty"`
	Filter             string `json:"filter,omitempty"`
	FilterTags         string `json:"filterTags,omitempty"`
	MatchesUri         string `json:"matchesUri,omitempty"`
}

var (
	filtersUri = "/compute-ops/v1beta1/filters"
)

func (c *ComClient) GetFilters(filters []string, offset, limit string) (FilterList, error) {
	var (
		q          map[string]interface{}
		filterList FilterList
	)
	c.SetAuthHeaderOptions(c.GetAuthHeaders())
	q = c.SetQueryParams(filters, offset, limit)
	data, err := c.RestAPICall(rest.GET, filtersUri, nil, q)
	if err != nil {
		return filterList, err
	}
	if err := json.Unmarshal([]byte(data), &filterList); err != nil {
		return filterList, err
	}
	return filterList, nil
}
