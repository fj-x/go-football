package footballdataapi

type ResultSet struct {
	Count  int16  `json:"count"`
	First  string `json:"first"`
	Last   string `json:"last"`
	Played int16  `json:"played"`
}

type Filter struct {
	Season   string `json:"season"`
	Matchday string `json:"matchday"`
}

type Area struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type Season struct {
	Id              int32  `json:"id"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	CurrentMatchday int16  `json:"currentMatchday"`
}

type Competition struct {
	Id     int16  `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblem"`
}
