package com

type FiterList struct {
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
