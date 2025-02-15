package model

func Get(title string) CampaignType {
	campaign := CampaignType{Title: title, Id: 1}

	return campaign
}
