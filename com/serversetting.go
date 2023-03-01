package com

import (
	"encoding/json"

	"example.local/hpecom/rest"
)

type BaseResult struct {
	Offset int `json:"offset,omitempty"`
	Count  int `jsont:"count,omitempty"`
	Total  int `json:"total,omitempty"`
	// Items  []ServerSettings `json:"items,omitempty"`
}

type AllServerSettings struct {
	BaseResult
	Items []AllSettingItems `json:"items,omitempty"`
}

type BiosSettingList struct {
	BaseResult
	Items []BiosSettingItems `json:"items,omitempty"`
}

type FirmwareSettingList struct {
	BaseResult
	Items []FirmwareSettingItems `json:"items,omitempty"`
}

type StorageSettingList struct {
	BaseResult
	Items []StorageSettingItems `json:"items,omitempty"`
}

type SettingItem struct {
	Id             string     `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	Description    string     `json:"description,omitempty"`
	Category       string     `json:"category,omitempty"`
	PlatformFamily string     `json:"platformFamily,omitempty"`
	Generation     int        `json:"generation,omitempty"`
	ResourceUri    string     `json:"resourceUri,omitempty"`
	Type           string     `json:"type,omitempty"`
	CreatedAt      string     `json:"createdAt,omitempty"`
	UpdatedAt      string     `json:"updatedAt,omitempty"`
	Client         *ComClient `json:"-"`
}

type AllSettingItems struct {
	SettingItem
	Settings interface{} `json:"settings"`
}

type BiosSettingItems struct {
	SettingItem
	Settings BiosSettings `json:"settings"`
}

type FirmwareSettingItems struct {
	SettingItem
	Settings FirmwareSettings `json:"settings"`
}

type StorageSettingItems struct {
	SettingItem
	Settings StorageSettings `json:"settings"`
}

type FirmwareSettings struct {
	GEN10 Gen10Bundle `json:"GEN10,omitempty"`
}

type Gen10Bundle struct {
	Id string `json:"id,omitempty"`
}

type BiosSettings struct {
	GEN10 Gen10Bios `json:"GEN10,omitempty"`
	GEN11 Gen11Bios `json:"GEN11,omitempty"`
}

type Gen10Bios struct {
	Attributes Attributes `json:"Attributes,omitempty"`
}

type Gen11Bios struct {
	Attributes Attributes `json:"Attributes,omitempty"`
}

type Attributes struct {
	WorkloadProfile string `json:"WorkloadProfile,omitempty"`
}

type StorageSettings struct {
	DefaultSettings DefaultSettings `json:"DEFAULT,omitempty"`
}

type DefaultSettings struct {
	RaidType       string `json:"raidType,omitempty"`
	VolumeSizeInGB int    `json:"volumeSizeInGB,omitempty"`
}

var (
	settingUri = "/compute-ops/v1beta1/server-settings"
	q          map[string]interface{}
)

func (c *ComClient) GetServerSettings(filters []string, offset, limit string) (AllServerSettings, error) {

	var (
		settingsList AllServerSettings
	)

	c.SetAuthHeaderOptions(c.GetAuthHeaders())

	q = c.SetQueryParams(filters, offset, limit)

	data, err := c.RestAPICall(rest.GET, settingUri, nil, q)
	if err != nil {
		return settingsList, err
	}

	if err := json.Unmarshal([]byte(data), &settingsList); err != nil {
		return settingsList, err
	}

	for i := 0; i < settingsList.Count; i++ {
		settingsList.Items[i].Client = c
	}

	return settingsList, nil
}

func (c *ComClient) GetBiosServerSettings() (BiosSettingList, error) {

	var biosSettingsList BiosSettingList
	query := []string{}
	query = append(query, "category eq 'BIOS'")

	thisSettings, err := c.GetServerSettings(query, "", "")
	if err != nil {
		return biosSettingsList, err
	}

	b, err := json.Marshal(thisSettings)
	if err != nil {
		return biosSettingsList, err
	}

	if err := json.Unmarshal([]byte(b), &biosSettingsList); err != nil {
		return biosSettingsList, err
	}

	return biosSettingsList, nil
}

func (c *ComClient) GetFirmwareServerSettings() (FirmwareSettingList, error) {

	var firmwareSettingsList FirmwareSettingList
	query := []string{}
	query = append(query, "category eq 'FIRMWARE'")

	thisSettings, err := c.GetServerSettings(query, "", "")
	if err != nil {
		return firmwareSettingsList, err
	}

	b, err := json.Marshal(thisSettings)
	if err != nil {
		return firmwareSettingsList, err
	}

	if err := json.Unmarshal([]byte(b), &firmwareSettingsList); err != nil {
		return firmwareSettingsList, err
	}

	return firmwareSettingsList, err
}

func (c *ComClient) GetStorageServerSettings() (StorageSettingList, error) {

	var storageSettingsList StorageSettingList
	query := []string{}
	query = append(query, "category eq 'STORAGE'")

	thisSettings, err := c.GetServerSettings(query, "", "")
	if err != nil {
		return storageSettingsList, err
	}

	b, err := json.Marshal(thisSettings)
	if err != nil {
		return storageSettingsList, err
	}

	if err := json.Unmarshal([]byte(b), &storageSettingsList); err != nil {
		return storageSettingsList, err
	}

	return storageSettingsList, err
}
