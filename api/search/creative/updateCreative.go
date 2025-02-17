package creative

import (
	"github.com/564104865/baidu-marketing/core"
	"github.com/564104865/baidu-marketing/model"
	"github.com/564104865/baidu-marketing/model/search/creative"
)

// UpdateCreative 修改推广创意
func UpdateCreative(clt *core.SDKClient, auth model.RequestHeader, creatives []creative.Creative) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body: creative.UpdateCreativeRequest{
			CreativeTypes: creatives,
		},
	}
	var resp creative.UpdateCreativeResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
