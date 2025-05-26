package route

// AddRouteStationReqDTO defines the request for adding a station to a route
type AddRouteStationReqDTO struct {
	RouteUuid         string  `json:"route_uuid" v:"required#线路UUID不能为空" dc:"线路UUID"`
	StationUuid       string  `json:"station_uuid" v:"required#站点UUID不能为空" dc:"站点UUID"`
	Sequence          int     `json:"sequence" v:"required#站点顺序不能为空" dc:"站点顺序"`
	DistanceFromStart float64 `json:"distance_from_start" v:"required#距起点距离不能为空" dc:"距起点距离(km)"`
	EstimatedTime     int     `json:"estimated_time" v:"required#预计到达时间不能为空" dc:"预计到达时间(分钟)"`
}

// UpdateRouteStationReqDTO defines the request for updating a station in a route
type UpdateRouteStationReqDTO struct {
	RouteStationUuid  string  `json:"route_station_uuid" v:"required#线路站点UUID不能为空" dc:"线路站点UUID"`
	Sequence          int     `json:"sequence" v:"required#站点顺序不能为空" dc:"站点顺序"`
	DistanceFromStart float64 `json:"distance_from_start" v:"required#距起点距离不能为空" dc:"距起点距离(km)"`
	EstimatedTime     int     `json:"estimated_time" v:"required#预计到达时间不能为空" dc:"预计到达时间(分钟)"`
}

// DeleteRouteStationReqDTO defines the request for deleting a station from a route
type DeleteRouteStationReqDTO struct {
	RouteStationUuid string `json:"route_station_uuid" v:"required#线路站点UUID不能为空" dc:"线路站点UUID"`
}

// GetRouteStationsReqDTO defines the request for getting stations in a route
type GetRouteStationsReqDTO struct {
	RouteUuid string `json:"route_uuid" v:"required#线路UUID不能为空" dc:"线路UUID"`
}

// RouteStationsDTO defines a list of stations in a route
type RouteStationsDTO struct {
	RouteUuid string                 `json:"route_uuid" dc:"线路UUID"`
	Stations  []*RouteStationItemDTO `json:"stations" dc:"站点列表"`
}

// RouteStationItemDTO defines a single station in a route
type RouteStationItemDTO struct {
	RouteStationUuid  string  `json:"route_station_uuid" dc:"线路站点UUID"`
	StationUuid       string  `json:"station_uuid" dc:"站点UUID"`
	Name              string  `json:"name" dc:"站点名称"`
	Code              string  `json:"code" dc:"站点编码"`
	Address           string  `json:"address" dc:"站点地址"`
	Longitude         float64 `json:"longitude" dc:"经度"`
	Latitude          float64 `json:"latitude" dc:"纬度"`
	Sequence          int     `json:"sequence" dc:"站点顺序"`
	DistanceFromStart float64 `json:"distance_from_start" dc:"距起点距离(km)"`
	EstimatedTime     int     `json:"estimated_time" dc:"预计到达时间(分钟)"`
}
