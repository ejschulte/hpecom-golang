package com

import (
	"encoding/json"
	"time"

	"github.com/ejschulte/hpecom-golang/rest"
)

type ScheduleList struct {
	Offset int        `json:"offset,omitempty"`
	Count  int        `jsont:"count,omitempty"`
	Total  int        `json:"total,omitempty"`
	Items  []Schedule `json:"items,omitempty"`
}

type Schedule struct {
	Id                    string            `json:"id,omitempty"`
	Type                  string            `json:"type,omitempty"`
	Generation            int               `json:"generation,omitempty"`
	CreatedAt             string            `json:"createdAt,omitempty"`
	UpdatedAt             string            `json:"updatedAt,omitempty"`
	ResourceUri           string            `json:"resourceUri,omitempty"`
	HistoryUri            string            `json:"historyUri,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Description           string            `json:"description,omitempty"`
	Purpose               string            `json:"purpose,omitempty"`
	AssociatedResourceUri string            `json:"associatedResourceUri,omitempty"`
	Schedule              ScheduleDetails   `json:"schedule,omitempty"`
	NextStartAt           string            `json:"nextStartAt,omitempty"`
	Operation             ScheduleOperation `json:"operation,omitempty"`
	// LastRun               ScheduleLastRun `json:"lastRun,omitempty"`
}

type ScheduleDetails struct {
	StartAt  string `json:"startAt,omitempty"`
	Interval string `json:"interval,omitempty"`
}

type ScheduleOperation struct {
	Type         string      `json:"type,omitempty"`
	TimeoutInSec float64     `json:"timeoutInSec,omitempty"`
	Method       string      `json:"method,omitempty"`
	Uri          string      `json:"uri,omitempty"`
	Query        interface{} `json:"query,omitempty"`
	Headers      interface{} `json:"headers,omitempty"`
	Body         ScheduleBody
}

type ScheduleBody struct {
	Data           ScheduleData `json:"data,omitempty"`
	ResourceUri    string       `json:"resourceUri,omitempty"`
	JobTemplateUri string       `json:"jobTemplateUri,omitempty"`
}

type ScheduleData struct {
}

type ScheduleLastRun struct {
	ID            string    `json:"id,omitempty"`
	OperationType string    `json:"operationType,omitempty"`
	DebugID       string    `json:"debugId,omitempty"`
	StartedAt     time.Time `json:"startedAt,omitempty"`
	DurationInSec float64   `json:"durationInSec,omitempty"`
	Summary       string    `json:"summary,omitempty"`
	Succeeded     bool      `json:"succeeded,omitempty"`
	ScheduleID    string    `json:"scheduleId,omitempty"`
	ScheduleURI   string    `json:"scheduleUri,omitempty"`
	ResourceURI   string    `json:"resourceUri,omitempty"`
	Type          string    `json:"type,omitempty"`
	Generation    int       `json:"generation,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	Status        int       `json:"status,omitempty"`
	Headers       struct {
		Date          string `json:"Date,omitempty"`
		Server        string `json:"Server,omitempty"`
		ContentType   string `json:"Content-Type,omitempty"`
		ContentLength string `json:"Content-Length,omitempty"`
	} `json:"headers,omitempty"`
	Body struct {
		ID   string `json:"id,omitempty"`
		Data struct {
		} `json:"data,omitempty,omitempty"`
		Name    string `json:"name,omitempty"`
		Type    string `json:"type,omitempty"`
		State   string `json:"state,omitempty"`
		Status  string `json:"status,omitempty"`
		Results struct {
		} `json:"results,omitempty"`
		SelfURI  string `json:"selfUri,omitempty"`
		OriginID string `json:"originId,omitempty"`
		Resource struct {
			Type        string `json:"type,omitempty"`
			ResourceURI string `json:"resourceUri,omitempty"`
		} `json:"resource,omitempty"`
		CreatedAt             string `json:"createdAt,omitempty"`
		UpdatedAt             string `json:"updatedAt,omitempty"`
		Generation            int    `json:"generation,omitempty"`
		ModifiedAt            string `json:"modifiedAt,omitempty"`
		DisplayName           string `json:"displayName,omitempty"`
		ParentJobID           string `json:"parentJobId,omitempty"`
		ResourceURI           string `json:"resourceUri,omitempty"`
		StatusDetails         string `json:"statusDetails,omitempty"`
		JobTemplateURI        string `json:"jobTemplateUri,omitempty"`
		AssociatedResourceURI string `json:"associated_resource_uri,omitempty"`
	} `json:"body,omitempty"`
}

var (
	scheduleUri = "/compute-ops/v1beta2/schedules"
)

func (c *ComClient) GetSchedules(filters []string, offset, limit string) (ScheduleList, error) {

	var (
		q         map[string]interface{}
		schedules ScheduleList
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaders())
	q = c.SetQueryParams(filters, offset, limit)

	data, err := c.RestAPICall(rest.GET, scheduleUri, nil, q)
	if err != nil {
		return schedules, err
	}

	if err := json.Unmarshal([]byte(data), &schedules); err != nil {
		return schedules, err
	}

	return schedules, nil
}
