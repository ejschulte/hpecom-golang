package com

import (
	"encoding/json"
	"fmt"

	"github.com/ejschulte/hpecom-golang/rest"
)

type ServerList struct {
	Offset int      `json:"offset,omitempty"`
	Count  int      `jsont:"count,omitempty"`
	Total  int      `json:"total,omitempty"`
	Items  []Server `json:"items,omitempty"`
}
type Server struct {
	Id                 string             `json:"id,omitempty"`
	Type               string             `json:"type,omitempty"`
	PlatformFamily     string             `json:"platformFamily,omitempty"`
	DeviceClaimType    string             `json:"deviceClaimType,omitempty"`
	ResourceURI        string             `json:"resourceUri,omitempty"`
	Name               string             `json:"name,omitempty"`
	CreatedAt          string             `json:"createdAt,omitempty"`
	UpdatedAt          string             `json:"updatedAt,omitempty"`
	Generation         int                `json:"generation,omitempty"`
	Hardware           ServerHardware     `json:"hardware,omitempty"`
	State              ServerState        `json:"state,omitempty"`
	FirmwareInventory  []ServerFirmware   `json:"firmwareInventory,omitempty"`
	FirmwareBundleUri  string             `json:"firmwareBundleUri,omitempty"`
	LastFirmwareUpdate LastFirmwareUpdate `json:"lastFirmwareUpdate,omitempty"`
	Host               ServerHost         `json:"host,omitempty"`
	ProcessorVendor    string             `json:"processorVendor,omitempty"`
	BiosFamily         string             `json:"boisFamily,omitempty"`
	// Client             *ComClient         `json:"-"`
}

type ServerHardware struct {
	SerialNumber string       `json:"serialNumber,omitempty"`
	Model        string       `json:"model,omitempty"`
	UUID         string       `json:"uuid,omitempty"`
	ProductID    string       `json:"productId,omitempty"`
	PowerState   string       `json:"powerState,omitempty"`
	IndicatorLed string       `json:"indicatorLed,omitempty"`
	Health       ServerHealth `json:"health,omitempty"`
	BMC          ServerBMC    `json:"bmc,omitempty"`
}

type ServerHealth struct {
	Summary                 string `json:"summary,omitempty"`
	Fans                    string `json:"fans,omitempty"`
	FanRedundancy           string `json:"fanRedundancy,omitempty"`
	LiquidCooling           string `json:"liquidCooling,omitempty"`
	LiquidCoolingRedundancy string `json:"liquidCoolingRedundancy,omitempty"`
	Memory                  string `json:"memory,omitempty"`
	Network                 string `json:"network,omitempty"`
	PowerSupplies           string `json:"powerSupplies,omitempty"`
	PowerSupplyRedundancy   string `json:"powerSupplyRedundancy,omitempty"`
	Processor               string `json:"processor,omitempty"`
	Storage                 string `json:"storage,omitempty"`
	Temperature             string `json:"temperature,omitempty"`
	Bios                    string `json:"bios,omitempty"`
	SmartStorage            string `json:"smartStorage,omitempty"`
	HealthLED               string `json:"healthLED,omitempty"`
}

type ServerBMC struct {
	Mac      string `json:"mac,omitempty"`
	IP       string `json:"ip,omitempty"`
	HostName string `json:"hostname,omitempty"`
}

type ServerState struct {
	Managed               bool   `json:"managed,omitempty"`
	Connected             bool   `json:"connected"`
	SubscriptionState     string `json:"subscriptionState,omitempty"`
	SubscriptionTier      string `json:"subscriptionTier,omitempty"`
	SubscriptionExpiresAt string `json:"subscriptionExpiresAt,omitempty"`
}

type ServerFirmware struct {
	Name          string `json:"name,omitempty"`
	Version       string `json:"version,omitempty"`
	DeviceContext string `json:"deviceContext,omitempty"`
}

type LastFirmwareUpdate struct {
	Status                   string            `json:"status,omitempty"`
	AttemptedAt              string            `json:"attempteAt,omitempty"`
	FirmwareInventoryUpdates []FirmwareUpdates `json:"firmwareInventoryUpdates,omitempty"`
}

type FirmwareUpdates struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Status  string `json:"status,omitempty"`
}

type ServerHost struct {
	OsName        string `json:"osName,omitempty"`
	OsVersion     string `json:"osVersion,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	OsType        string `json:"osType,omitempty"`
	OsDescription string `json:"osDescription,omitempty"`
}

var (
	serversUri = "/compute-ops/v1beta2/servers"
)

func (c *ComClient) GetServers(filters []string, offset, limit string) (ServerList, error) {

	var (
		q       map[string]interface{}
		servers ServerList
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaders())

	q = c.SetQueryParams(filters, offset, limit)

	data, err := c.RestAPICall(rest.GET, serversUri, nil, q)
	if err != nil {
		return servers, err
	}

	if err := json.Unmarshal([]byte(data), &servers); err != nil {
		return servers, err
	}

	return servers, nil
}

func (c *ComClient) GetServerByName(name string) (Server, error) {

	query := []string{}
	query = append(query, fmt.Sprintf("contains(host/hostname,'%s')", name))

	servers, err := c.GetServers(query, "", "")

	if servers.Total > 0 {
		return servers.Items[0], err
	}

	return Server{}, err
}
