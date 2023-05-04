package endpoints

import (
	"emailin/internal/infrastructure/database"
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	repository := &database.CampaignRepository{}

	campaigns, err := repository.Get()

	if err != nil {
		return nil, 400, err
	}

	return campaigns, 200, nil
}
