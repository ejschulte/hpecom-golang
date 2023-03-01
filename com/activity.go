package com

import "time"

type ActivityList struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Generation  int       `json:"generation"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	ResourceURI string    `json:"resourceUri"`
	Source      struct {
		Type        string      `json:"type"`
		ResourceURI string      `json:"resourceUri"`
		DisplayName string      `json:"displayName"`
		OriginID    interface{} `json:"originId"`
	} `json:"source"`
	Title                  string      `json:"title"`
	Key                    string      `json:"key"`
	Message                string      `json:"message"`
	CreatedAt              time.Time   `json:"createdAt"`
	RecommendedAction      interface{} `json:"recommendedAction"`
	HealthState            string      `json:"healthState"`
	AssociatedServerURI    string      `json:"associatedServerUri"`
	AssociatedServerID     string      `json:"associatedServerId"`
	GroupID                string      `json:"groupId"`
	GroupDisplayName       string      `json:"groupDisplayName"`
	ServerSettingsID       string      `json:"serverSettingsId"`
	ServerSettingsName     string      `json:"serverSettingsName"`
	ServerSettingsCategory string      `json:"serverSettingsCategory"`
}
