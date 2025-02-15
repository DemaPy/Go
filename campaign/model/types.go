package model

type CampaignType struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ResponseType struct {
	Data    any    `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
