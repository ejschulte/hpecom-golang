package com

import "time"

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
	TimeoutInSec string      `json:"timeoutInSec,omitempty"`
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
	ID            string    `json:"id"`
	OperationType string    `json:"operationType"`
	DebugID       string    `json:"debugId"`
	StartedAt     time.Time `json:"startedAt"`
	DurationInSec float64   `json:"durationInSec"`
	Summary       string    `json:"summary"`
	Succeeded     bool      `json:"succeeded"`
	ScheduleID    string    `json:"scheduleId"`
	ScheduleURI   string    `json:"scheduleUri"`
	ResourceURI   string    `json:"resourceUri"`
	Type          string    `json:"type"`
	Generation    int       `json:"generation"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Status        int       `json:"status"`
	Headers       struct {
		Date          string `json:"Date"`
		Server        string `json:"Server"`
		ContentType   string `json:"Content-Type"`
		ContentLength string `json:"Content-Length"`
	} `json:"headers"`
	Body struct {
		ID   string `json:"id"`
		Data struct {
		} `json:"data,omitempty"`
		Name    string `json:"name"`
		Type    string `json:"type"`
		State   string `json:"state"`
		Status  string `json:"status"`
		Results struct {
		} `json:"results"`
		SelfURI  string `json:"selfUri"`
		OriginID string `json:"originId"`
		Resource struct {
			Type        string `json:"type"`
			ResourceURI string `json:"resourceUri"`
		} `json:"resource"`
		CreatedAt             string `json:"createdAt"`
		UpdatedAt             string `json:"updatedAt"`
		Generation            int    `json:"generation"`
		ModifiedAt            string `json:"modifiedAt"`
		DisplayName           string `json:"displayName"`
		ParentJobID           string `json:"parentJobId,omitempty"`
		ResourceURI           string `json:"resourceUri"`
		StatusDetails         string `json:"statusDetails,omitempty"`
		JobTemplateURI        string `json:"jobTemplateUri"`
		AssociatedResourceURI string `json:"associated_resource_uri"`
	} `json:"body"`
}

var (
	scheduleUri = "/compute-ops/v1beta2/schedules"
)
