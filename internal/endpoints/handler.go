package endpoints

import "emailin/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
