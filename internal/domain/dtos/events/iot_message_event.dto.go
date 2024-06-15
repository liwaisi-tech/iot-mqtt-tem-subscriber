package dtos

type IOTMessageEventDTO struct {
	MACAddress  string  `json:"mac_address"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
