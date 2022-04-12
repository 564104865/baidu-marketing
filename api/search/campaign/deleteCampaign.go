package campaign

import (
	"github.com/564104865/baidu-marketing/core"
	"github.com/564104865/baidu-marketing/model"
	"github.com/564104865/baidu-marketing/model/search/campaign"
)

// DeleteCampaign 删除计划
func DeleteCampaign(clt *core.SDKClient, auth model.RequestHeader, campaignIds []int64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: campaign.DeleteCampaignRequest{
			CampaignIds: campaignIds,
		},
	}

	return clt.Do(req, nil)
}
