package campaign

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type AddCampaignRequest struct {
	CampaignTypes []Campaign `json:"campaignTypes,omitempty"` // 新增推广计划物料;集合长度限制：[1, 100]
	AdType        int        `json:"adType,omitempty"`        // 投放广告类型
}

func (r AddCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/addCampaign", model.BASE_URL_SMS)
}