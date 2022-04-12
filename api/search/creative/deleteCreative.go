package creative

import (
	"github.com/564104865/baidu-marketing/core"
	"github.com/564104865/baidu-marketing/model"
	"github.com/564104865/baidu-marketing/model/search/creative"
)

// DeleteCreative 删除推广创意
func DeleteCreative(clt *core.SDKClient, auth model.RequestHeader, creativeIds []int64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: creative.DeleteCreativeRequest{
			CreativeIds: creativeIds,
		},
	}

	return clt.Do(req, nil)
}
