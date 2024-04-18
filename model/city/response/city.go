package response

type GaodeResp struct {
	Key      string `json:"key"`
	Keywords string `json:"keywords"`
	Types    string `json:"types"`
	City     string `json:"city"`
	Children int    `json:"children"`
	Offset   int    `json:"offset"`
	Page     int    `json:"page"`
}

type GetLocalResp struct {
	Type     string `json:"type"`
	CityID   int64  `json:"cityID"`
	CityName string `json:"cityName"`
	Poi      []Poi  `json:"poi"`
}

type Poi struct {
	AreaID   int64  `json:"areaID"`
	AreaName string `json:"areaName"`
	Location string `json:"location"`
}
