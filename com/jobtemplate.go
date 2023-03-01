package com

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ejschulte/hpecom-golang/rest"
)

type JobTemplateList struct {
	Offset int           `json:"offset,omitempty"`
	Count  int           `jsont:"count,omitempty"`
	Total  int           `json:"total,omitempty"`
	Items  []JobTemplate `json:"items,omitempty"`
}

type JobTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	ResourceURI string `json:"resourceUri"`
	Schema      struct {
		Type       string `json:"type,omitempty"`
		Properties struct {
		} `json:"properties,omitempty"`
	} `json:"schema,omitempty"`
	Type       string    `json:"type"`
	Generation int       `json:"generation"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// Create a Method type
type Template int

const (
	PowerOnNew Template = 1 + iota
	SppUpdate
	GroupFirmwareCompliance
	GetSettingsForTemplate
	RestartNew
	ColdBoot
	PowerOffNew
	FirmwareUpdateNew
	GroupAddActions
	IloOnlyFirmwareUpdate
	GroupFirmwareUpdate
	DataRoundupReportOrchestrator
	ApplySettingsTemplate
	GroupStorageVolumeConfiguration
	GroupApplyServerSettings
	GroupOSInstallation
	CreateOneviewComplianceReport
)

var template = [...]string{
	"0cbb2377-1834-488d-840c-d5bf788c34fb",
	"2626e682-7bd4-40e3-b93a-745133deecd2",
	"23b8ba2a-6c46-4223-b028-919382c7dcac",
	"6cd671db-ce6b-45ce-894e-7b5ae23e0399",
	"30110551-cad6-4069-95b8-dbce9bbd8525",
	"aacfb3e0-6575-4d4f-a711-1ee1ae768407",
	"d0c13b58-748c-461f-9a61-c0c5c71f1bb4",
	"fd54a96c-cabc-42e3-aee3-374a2d009dba",
	"7983f2de-1ea7-4399-91ec-62d3647409e0",
	"94caa4ef-9ff8-4805-9e97-18a09e673b66",
	"91159b5e-9eeb-11ec-a9da-00155dc0a0c0",
	"b0001d36-6490-48ac-93af-a87adfb997ed",
	"2d0f40f7-2a07-4c74-92e1-d1afaf49e632",
	"c708eb57-235d-4ea8-9e21-8ceea2438773",
	"beff07ce-f36d-4699-9ac3-f872dcd63133",
	"e2952628-2629-4088-93db-91742304ef0c",
	"aae145a1-79a2-4516-b191-c98039c96542",
}

func (t Template) String() string {
	return template[t-1]
}

var (
	jobTemplatesUri = "/compute-ops/v1beta2/job-templates"
)

func (c *ComClient) GetJobTemplates(filters []string, offset, limit string) (JobTemplateList, error) {
	var (
		q            map[string]interface{}
		jobTemplates JobTemplateList
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaders())

	q = c.SetQueryParams(filters, offset, limit)

	data, err := c.RestAPICall(rest.GET, jobTemplatesUri, nil, q)
	if err != nil {
		return jobTemplates, err
	}

	if err := json.Unmarshal([]byte(data), &jobTemplates); err != nil {
		return jobTemplates, err
	}

	return jobTemplates, nil
}

func (c *ComClient) GetJobTemplateByName(t Template) (JobTemplate, error) {

	var (
		jobTemplate JobTemplate
	)

	id := t.String()
	uri := jobTemplatesUri + "/" + id

	fmt.Println(uri)

	c.SetAuthHeaderOptions(c.GetAuthHeaders())

	data, err := c.RestAPICall(rest.GET, uri, nil, nil)
	if err != nil {
		return jobTemplate, err
	}

	if err := json.Unmarshal([]byte(data), &jobTemplate); err != nil {
		return jobTemplate, err
	}

	return jobTemplate, nil

}
