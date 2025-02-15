package utils

import (
	"encoding/json"
	"go_dev/campaign/model"
)

func ResponseDataError(err any) model.ResponseType {
	return model.ResponseType{
		Success: false,
		Message: "Failed",
		Data:    err,
	}
}

func ResponseData(payload model.CampaignType) model.ResponseType {
	return model.ResponseType{
		Success: true,
		Message: "Success",
		Data:    model.CampaignType{Id: payload.Id, Title: payload.Title},
	}
}

func ToJson(response any) ([]byte, error) {
	return json.Marshal(response)
}

func ResponseDataToJson(payload model.CampaignType) ([]byte, error) {
	response := model.ResponseType{
		Success: true,
		Message: "Success",
		Data:    model.CampaignType{Id: payload.Id, Title: payload.Title},
	}
	return json.Marshal(response)
}
