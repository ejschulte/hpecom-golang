package com

import (
	"encoding/json"
	"time"

	"example.local/hpecom/rest"
)

type ReportList struct {
	Offset int      `json:"offset,omitempty"`
	Count  int      `jsont:"count,omitempty"`
	Total  int      `json:"total,omitempty"`
	Items  []Report `json:"items,omitempty"`
}

type Report struct {
	Id                    string    `json:"id,omitempty"`
	Type                  string    `json:"type,omitempty"`
	Generation            int       `json:"generation,omitempty"`
	CreatedAt             time.Time `json:"createdAt,omitempty"`
	UpdatedAt             time.Time `json:"updatedAt,omitempty"`
	ResourceUri           string    `json:"resourceUri,omitempty"`
	Name                  string    `json:"name,omitempty"`
	ReportDataStartAt     time.Time `json:"reportDataStartAt,omitempty"`
	ReportDataEndAt       time.Time `json:"reportDataEndAt,omitempty"`
	ReportDataUri         string    `json:"reportDataUri,omitempty"`
	ReportType            string    `json:"reportType,omitempty"`
	ReportTypeDisplayName string    `json:"reportTypeDisplayName,omitempty"`
	DeviceIds             []string  `json:"deviceIds,omitempty"`
}

var (
	reportsUri = "/compute-ops/v1beta1/reports"
)

func (c *ComClient) GetReports(filters []string, offset, limit string) (ReportList, error) {
	var (
		q       map[string]interface{}
		reports ReportList
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaders())

	q = c.SetQueryParams(filters, offset, limit)

	data, err := c.RestAPICall(rest.GET, reportsUri, nil, q)
	if err != nil {
		return reports, err
	}

	if err := json.Unmarshal([]byte(data), &reports); err != nil {
		return reports, err
	}

	return reports, nil
}
