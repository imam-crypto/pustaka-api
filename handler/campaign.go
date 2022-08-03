package handler

import (
	"net/http"
	"pustaka-api/campaign"
	"pustaka-api/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.FindByUserID(userID)

	if err != nil {

		response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data Not Found", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
}
