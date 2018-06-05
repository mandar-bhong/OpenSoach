package models

type APIDashboardDeviceSummaryResponse struct {
	TotalDevices   int `json:"total"`
	Onlinedevices  int `json:"online"`
	Offlinedevices int `json:"offline"`
}

type APIDashboardLocationSummaryResponse struct {
	Total  int `json:"total"`
	Active int `json:"active"`
}
