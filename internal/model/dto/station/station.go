package station

// StationDTO defines a station
type StationDTO struct {
	StationUuid string  `json:"station_uuid" dc:"站点UUID"`
	Name        string  `json:"name" dc:"站点名称"`
	Code        string  `json:"code" dc:"站点编码"`
	Address     string  `json:"address" dc:"站点地址"`
	Longitude   float64 `json:"longitude" dc:"经度"`
	Latitude    float64 `json:"latitude" dc:"纬度"`
	Status      int     `json:"status" dc:"状态: 0-停用, 1-启用"`
	CreatedAt   string  `json:"created_at" dc:"创建时间"`
	UpdatedAt   string  `json:"updated_at" dc:"更新时间"`
}

// StationsListDTO defines a list of stations
type StationsListDTO struct {
	Total    int          `json:"total" dc:"总数"`
	Page     int          `json:"page" dc:"页码"`
	Size     int          `json:"size" dc:"每页数量"`
	Stations []*StationDTO `json:"stations" dc:"站点列表"`
}

// StationListReqDTO defines the request for listing stations
type StationListReqDTO struct {
	Name   string `json:"name" dc:"站点名称，用于模糊查询"`
	Code   string `json:"code" dc:"站点编码，用于模糊查询"`
	Status int    `json:"status" dc:"状态: 0-停用, 1-启用"`
	Page   int    `json:"page" dc:"页码，默认为1"`
	Size   int    `json:"size" dc:"每页数量，默认为10"`
}