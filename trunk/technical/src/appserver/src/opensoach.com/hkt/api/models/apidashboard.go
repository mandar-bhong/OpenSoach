package models

type APIDashboardDeviceSummaryResponse struct {
	TotalDevices   int `json:"totaldevices"`
	Onlinedevices  int `json:"onlinedevices"`
	Offlinedevices int `json:"offlinedevices"`
}
